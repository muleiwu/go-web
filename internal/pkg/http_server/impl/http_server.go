package impl

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	env        interfaces.EnvInterface
	logger     interfaces.LoggerInterface
	middleware []gin.HandlerFunc
	routerFunc func(router *gin.Engine)
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

// 启动HTTP服务器并注册路由和中间件
func (receiver *HttpServer) runHttp() {
	// 设置Gin模式
	if receiver.env.GetString("mode", "") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 完全替换gin的默认Logger
	gin.DisableConsoleColor()
	//zapLogger := receiver.logger
	//gin.DefaultWriter = &zapLogWriter{zapLogger: zapLogger}
	//gin.DefaultErrorWriter = &zapLogWriter{zapLogger: zapLogger, isError: true}

	// 配置Gin引擎
	// 配置Gin引擎并替换默认logger
	engine := gin.New()
	engine.Use(gin.Recovery())
	//engine.Use(GinZapLogger())

	// 注册中间件
	//handlerFuncs := config.MiddlewareConfig{}.Get()
	for i, handlerFunc := range receiver.middleware {
		if handlerFunc == nil {
			continue
		}
		engine.Use(handlerFunc)
		receiver.logger.Info(fmt.Sprintf("注册中间件: %d", i))
	}

	//router.InitRouter(engine)
	receiver.routerFunc(engine)

	// 创建一个HTTP服务器，以便能够优雅关闭
	addr := receiver.env.GetString("addr", ":8080")
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	// 创建一个通道来接收中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 在单独的goroutine中启动服务器
	go func() {
		receiver.logger.Info(fmt.Sprintf("服务器启动于 %s", addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			receiver.logger.Error(fmt.Sprintf("启动服务器失败: %v", err))
		}
	}()

	// 等待中断信号
	<-quit
	receiver.logger.Info("正在关闭服务器...")

	// 创建一个5秒的上下文用于超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅地关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		receiver.logger.Error(fmt.Sprintf("服务器强制关闭: %v", err))
	}

	receiver.logger.Info("服务器已优雅关闭")
}
