package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mliev.com/template/go-web/config"
	"mliev.com/template/go-web/helper"
	"mliev.com/template/go-web/router"
)

// Start 启动应用程序
func Start() {
	initializeServices()
	go RunHttp()
	// 添加阻塞以保持主程序运行
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}

// initializeServices 初始化所有服务
func initializeServices() {
	if err := helper.InitViper(); err != nil {
		helper.Logger().Error(fmt.Sprintf("配置初始化失败: %v", err))
		os.Exit(1)
	}
	helper.GetDB()
	helper.GetRedis()
}

// zapLogWriter 实现io.Writer接口，将gin的日志输出重定向到zap
type zapLogWriter struct {
	zapLogger *zap.Logger
	isError   bool
}

// Write 实现io.Writer接口
func (z *zapLogWriter) Write(p []byte) (n int, err error) {
	if z.isError {
		z.zapLogger.Error(string(p))
	} else {
		z.zapLogger.Info(string(p))
	}
	return len(p), nil
}

// RunHttp 启动HTTP服务器并注册路由和中间件
func RunHttp() {

	// 设置Gin模式
	if helper.EnvString("mode", "") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 完全替换gin的默认Logger
	gin.DisableConsoleColor()
	zapLogger := helper.Logger()
	gin.DefaultWriter = &zapLogWriter{zapLogger: zapLogger}
	gin.DefaultErrorWriter = &zapLogWriter{zapLogger: zapLogger, isError: true}

	// 配置Gin引擎
	// 配置Gin引擎并替换默认logger
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(GinZapLogger())

	// 注册中间件
	handlerFuncs := config.MiddlewareConfig{}.Get()
	for i, handlerFunc := range handlerFuncs {
		if handlerFunc == nil {
			continue
		}
		engine.Use(handlerFunc)
		helper.Logger().Info(fmt.Sprintf("注册中间件: %d", i))
	}

	router.InitRouter(engine)

	// 创建一个HTTP服务器，以便能够优雅关闭
	addr := helper.EnvString("addr", ":8080")
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	// 创建一个通道来接收中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 在单独的goroutine中启动服务器
	go func() {
		helper.Logger().Info(fmt.Sprintf("服务器启动于 %s", addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			helper.Logger().Error(fmt.Sprintf("启动服务器失败: %v", err))
		}
	}()

	// 等待中断信号
	<-quit
	helper.Logger().Info("正在关闭服务器...")

	// 创建一个5秒的上下文用于超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅地关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		helper.Logger().Error(fmt.Sprintf("服务器强制关闭: %v", err))
	}

	helper.Logger().Info("服务器已优雅关闭")
}

// GinZapLogger 返回一个Gin中间件，使用zap记录HTTP请求
func GinZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		// 请求处理完成后记录日志
		cost := time.Since(start)
		zapLogger := helper.Logger()
		statusCode := c.Writer.Status()

		// 通用的日志字段
		fields := []zap.Field{
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.Int("status", statusCode),
			zap.String("ip", c.ClientIP()),
			zap.Duration("latency", cost),
			zap.String("user-agent", c.Request.UserAgent()),
		}

		// 根据状态码决定日志级别
		switch {
		case statusCode >= 500:
			fields = append(fields, zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()))
			zapLogger.Error("请求处理", fields...)
		case statusCode >= 400:
			zapLogger.Warn("请求处理", fields...)
		default:
			zapLogger.Info("请求处理", fields...)
		}
	}
}
