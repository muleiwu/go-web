package cmd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mliev.com/template/go-web/app/Middleware"
	"mliev.com/template/go-web/config"
	"mliev.com/template/go-web/router"
	"mliev.com/template/go-web/support"
	"mliev.com/template/go-web/support/db"
	"mliev.com/template/go-web/support/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {
	config.InitViper()
	db.GetDB()
	db.GetRedis()
	go RunHttp()
	// 添加阻塞以保持主程序运行
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
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

func RunHttp() {

	// 设置Gin模式
	if support.Env("mode", "").(string) == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 完全替换gin的默认Logger
	gin.DisableConsoleColor()
	zapLogger := logger.Get()
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
		logger.Get().Info(fmt.Sprintf("注册中间件: %d", i))
	}
	engine.Use(middleware.CorsMiddleware())
	router.InitRouter(engine)

	// 创建一个HTTP服务器，以便能够优雅关闭
	addr := support.Env("addr", ":8080").(string)
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	// 创建一个通道来接收中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 在单独的goroutine中启动服务器
	go func() {
		logger.Get().Info(fmt.Sprintf("服务器启动于 %s", addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Get().Error(fmt.Sprintf("启动服务器失败: %v", err))
		}
	}()

	// 等待中断信号
	<-quit
	logger.Get().Info("正在关闭服务器...")

	// 创建一个5秒的上下文用于超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅地关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		logger.Get().Error(fmt.Sprintf("服务器强制关闭: %v", err))
	}

	logger.Get().Info("服务器已优雅关闭")
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
		zapLogger := logger.Get()

		// 根据状态码决定日志级别
		statusCode := c.Writer.Status()
		if statusCode >= 500 {
			zapLogger.Error("请求处理",
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.Int("status", statusCode),
				zap.String("ip", c.ClientIP()),
				zap.Duration("latency", cost),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			)
		} else if statusCode >= 400 {
			zapLogger.Warn("请求处理",
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.Int("status", statusCode),
				zap.String("ip", c.ClientIP()),
				zap.Duration("latency", cost),
				zap.String("user-agent", c.Request.UserAgent()),
			)
		} else {
			zapLogger.Info("请求处理",
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.Int("status", statusCode),
				zap.String("ip", c.ClientIP()),
				zap.Duration("latency", cost),
				zap.String("user-agent", c.Request.UserAgent()),
			)
		}
	}
}
