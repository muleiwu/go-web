package impl

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"time"

	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/muleiwu/golog"
	"github.com/muleiwu/gsr"
)

type HttpServer struct {
	Helper     interfaces.HelperInterface
	routerFunc func(router *gin.Engine)
	server     *http.Server
}

func NewHttpServer(helper interfaces.HelperInterface) *HttpServer {
	return &HttpServer{
		Helper: helper,
	}
}

// RunHttp 启动HTTP服务器并注册路由和中间件
func (receiver *HttpServer) RunHttp() {
	// 设置Gin模式
	if receiver.Helper.GetConfig().GetString("http.mode", "") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 将 Gin 内部调试日志重定向到 golog
	gin.DisableConsoleColor()
	gin.DefaultWriter = &gologWriter{logger: receiver.Helper.GetLogger()}
	gin.DefaultErrorWriter = &gologWriter{logger: receiver.Helper.GetLogger(), isError: true}

	// 自定义路由注册日志，显示真实控制器方法名而非 WrapHandler.func1
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		if lastHandlerName != "" {
			handlerName = lastHandlerName
		}
		receiver.Helper.GetLogger().Info("路由注册",
			golog.Field("method", httpMethod),
			golog.Field("path", absolutePath),
			golog.Field("handler", handlerName),
			golog.Field("handlers", nuHandlers),
		)
	}

	// 配置Gin引擎并替换默认logger
	engine := gin.New()
	// 增加链路的追踪ID
	engine.Use(receiver.traceIdMiddleware())
	// 对路由的输入输出记录
	engine.Use(receiver.ginLogger())
	// 防止 panic 导致的程序崩溃
	engine.Use(gin.Recovery())

	// 加载HTML模板
	if err := receiver.loadTemplates(engine); err != nil {
		panic(err)
	}

	// 加载网站静态资源
	receiver.loadWebStatic(engine)

	// 注册中间件
	//handlerFuncs := config.MiddlewareConfig{}.Get()
	middlewareFuncList := receiver.Helper.GetConfig().Get("http.middleware", []gin.HandlerFunc{}).([]gin.HandlerFunc)
	for _, handlerFunc := range middlewareFuncList {
		if handlerFunc == nil {
			continue
		}
		engine.Use(handlerFunc)
		receiver.Helper.GetLogger().Info(fmt.Sprintf("注册中间件: %s", receiver.GetFunctionName(handlerFunc)))
	}
	receiver.Helper.GetLogger().Info(fmt.Sprintf("注册中间件: %d 个", len(middlewareFuncList)))

	deps := NewHttpDeps(receiver.Helper, engine)
	routerFunc := receiver.Helper.GetConfig().Get("http.router", func(r httpInterfaces.RouterInterface) {}).(func(httpInterfaces.RouterInterface))
	routerFunc(NewRouter(engine, deps))

	//receiver.routerFunc(engine)

	// 创建一个HTTP服务器，以便能够优雅关闭
	addr := receiver.Helper.GetConfig().GetString("http.addr", ":8080")
	receiver.server = &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	// 在单独的goroutine中启动服务器
	go func() {
		receiver.Helper.GetLogger().Info(fmt.Sprintf("服务器启动于 %s", addr))
		if err := receiver.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			receiver.Helper.GetLogger().Error(fmt.Sprintf("启动服务器失败: %v", err))
		}
	}()
}

// Stop 优雅停止HTTP服务器
func (receiver *HttpServer) Stop() error {
	if receiver.server == nil {
		return nil
	}

	receiver.Helper.GetLogger().Info("正在关闭HTTP服务器...")

	// 创建一个5秒的上下文用于超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅地关闭服务器
	if err := receiver.server.Shutdown(ctx); err != nil {
		receiver.Helper.GetLogger().Error(fmt.Sprintf("HTTP服务器关闭失败: %v", err))
		return err
	}

	receiver.Helper.GetLogger().Info("HTTP服务器已优雅关闭")
	return nil
}

// loadTemplates 加载模板到Gin引擎
func (receiver *HttpServer) loadTemplates(engine *gin.Engine) error {
	staticFs := receiver.Helper.GetConfig().Get("static.fs", map[string]embed.FS{}).(map[string]embed.FS)

	receiver.Helper.GetLogger().Info(fmt.Sprintf("加载静态文件系统, 键数量: %d", len(staticFs)))
	for key := range staticFs {
		receiver.Helper.GetLogger().Info(fmt.Sprintf("  - 静态文件系统键: %s", key))
	}

	templates, ok := staticFs["templates"]

	if !ok {
		return errors.New("没有模板目录需要初始化")
	}

	// 从嵌入的文件系统创建子文件系统
	subFS, err := fs.Sub(templates, "templates")
	if err != nil {
		return fmt.Errorf("创建子文件系统失败: %v", err)
	}

	// 收集所有 HTML 文件路径
	var templateFiles []string
	err = fs.WalkDir(subFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && len(path) > 5 && path[len(path)-5:] == ".html" {
			templateFiles = append(templateFiles, path)
			receiver.Helper.GetLogger().Info(fmt.Sprintf("  - 找到模板文件: %s", path))
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("遍历模板文件失败: %v", err)
	}

	if len(templateFiles) == 0 {
		receiver.Helper.GetLogger().Warn("没有找到任何模板文件")
		return nil
	}

	receiver.Helper.GetLogger().Info(fmt.Sprintf("共找到 %d 个模板文件", len(templateFiles)))

	// 创建模板实例，手动解析每个文件并保留完整路径作为模板名称
	tmpl := template.New("")
	for _, file := range templateFiles {
		// 读取模板文件内容
		content, err := fs.ReadFile(subFS, file)
		if err != nil {
			return fmt.Errorf("读取模板文件 %s 失败: %v", file, err)
		}

		// 使用完整路径作为模板名称
		_, err = tmpl.New(file).Parse(string(content))
		if err != nil {
			return fmt.Errorf("解析模板文件 %s 失败: %v", file, err)
		}

		receiver.Helper.GetLogger().Info(fmt.Sprintf("  - 解析模板: %s", file))
	}

	// 列出所有已定义的模板
	receiver.Helper.GetLogger().Info("已定义的模板:")
	for _, t := range tmpl.Templates() {
		receiver.Helper.GetLogger().Info(fmt.Sprintf("  - 模板名称: %s", t.Name()))
	}

	// 设置HTML模板
	engine.SetHTMLTemplate(tmpl)

	receiver.Helper.GetLogger().Info("模板加载成功")

	return nil
}

func (receiver *HttpServer) loadWebStatic(engine *gin.Engine) {
	staticHandler := NewStaticHandler(receiver.Helper, engine)
	staticHandler.setupStaticFileServers()
}

// GetFunctionName 获取函数名
func (receiver *HttpServer) GetFunctionName(i any) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (receiver *HttpServer) traceIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		newUUID, err := uuid.NewV7()
		if err != nil {
			newUUID = uuid.New()
		}
		c.Set("traceId", newUUID.String())
		c.Writer.Header().Set("trace-id", newUUID.String())
		c.Next()
	}
}

// gologWriter bridges io.Writer (used by gin.DefaultWriter / gin.DefaultErrorWriter)
// to gsr.Logger, so Gin's internal debug messages pass through golog.
type gologWriter struct {
	logger  gsr.Logger
	isError bool
}

func (w *gologWriter) Write(p []byte) (n int, err error) {
	msg := strings.TrimRight(string(p), "\n")
	if msg == "" {
		return len(p), nil
	}
	if w.isError {
		w.logger.Error(msg)
	} else {
		w.logger.Info(msg)
	}
	return len(p), nil
}

func (receiver *HttpServer) ginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		traceId := c.GetString("traceId")

		c.Next()

		// 请求处理完成后记录日志
		cost := time.Since(start)
		zapLogger := receiver.Helper.GetLogger()

		// 根据状态码决定日志级别
		statusCode := c.Writer.Status()
		if statusCode >= 500 {
			zapLogger.Error("请求处理",
				golog.Field("traceId", traceId),
				golog.Field("method", c.Request.Method),
				golog.Field("path", path),
				golog.Field("query", query),
				golog.Field("status", statusCode),
				golog.Field("ip", c.ClientIP()),
				golog.Field("latency", cost),
				golog.Field("user-agent", c.Request.UserAgent()),
				golog.Field("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			)
		} else if statusCode >= 400 {
			zapLogger.Warn("请求处理",
				golog.Field("traceId", traceId),
				golog.Field("method", c.Request.Method),
				golog.Field("path", path),
				golog.Field("query", query),
				golog.Field("status", statusCode),
				golog.Field("ip", c.ClientIP()),
				golog.Field("latency", cost),
				golog.Field("user-agent", c.Request.UserAgent()),
			)
		} else {
			zapLogger.Info("请求处理",
				golog.Field("traceId", traceId),
				golog.Field("method", c.Request.Method),
				golog.Field("path", path),
				golog.Field("query", query),
				golog.Field("status", statusCode),
				golog.Field("ip", c.ClientIP()),
				golog.Field("latency", cost),
				golog.Field("user-agent", c.Request.UserAgent()),
			)
		}
	}
}
