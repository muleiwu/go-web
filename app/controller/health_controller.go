package controller

import (
	"context"
	"time"

	"cnb.cool/mliev/open/go-web/app/constants"
	"cnb.cool/mliev/open/go-web/app/dto"
	"cnb.cool/mliev/open/go-web/pkg/container"
	httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type HealthController struct {
	BaseResponse
}

// GetHealth 健康检查接口
func (receiver HealthController) GetHealth(c httpInterfaces.RouterContextInterface) {
	healthStatus := dto.HealthStatus{
		Status:    "UP",
		Timestamp: time.Now().Unix(),
		Services:  make(map[string]any),
	}

	// 检查数据库连接
	dbStatus := receiver.checkDatabase()
	healthStatus.Services["database"] = dbStatus

	// 检查Redis连接
	redisStatus := receiver.checkRedis()
	healthStatus.Services["redis"] = redisStatus

	// 如果任何服务不健康，整体状态设为DOWN
	if dbStatus.Status == "DOWN" || redisStatus.Status == "DOWN" {
		healthStatus.Status = "DOWN"
		receiver.Error(c, constants.ErrCodeUnavailable, "服务不健康")
		return
	}

	receiver.Success(c, healthStatus)
}

// GetHealthSimple 简单健康检查接口
func (receiver HealthController) GetHealthSimple(c httpInterfaces.RouterContextInterface) {
	receiver.Success(c, map[string]any{
		"status":    "UP",
		"timestamp": time.Now().Unix(),
	})
}

// checkDatabase 检查数据库连接
func (receiver HealthController) checkDatabase() dto.ServiceStatus {
	// 不用 helper.GetDatabase()：它走 container.MustGet，数据库未注册时直接 panic，
	// 健康检查应把「未装配」报告为 DOWN 而不是 500。
	gormDB, getErr := container.Get[*gorm.DB]()
	if getErr != nil || gormDB == nil {
		return dto.ServiceStatus{
			Status:  "DOWN",
			Message: "数据库连接失败",
		}
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return dto.ServiceStatus{
			Status:  "DOWN",
			Message: "获取底层数据库连接失败: " + err.Error(),
		}
	}

	if err := sqlDB.Ping(); err != nil {
		return dto.ServiceStatus{
			Status:  "DOWN",
			Message: "数据库ping失败: " + err.Error(),
		}
	}

	return dto.ServiceStatus{
		Status: "UP",
	}
}

// checkRedis 检查Redis连接
func (receiver HealthController) checkRedis() dto.ServiceStatus {
	// 同 checkDatabase：不走 MustGet，未注册报 DOWN 而非 panic。
	redisClient, getErr := container.Get[*redis.Client]()
	if getErr != nil || redisClient == nil {
		return dto.ServiceStatus{
			Status:  "DOWN",
			Message: "Redis连接失败",
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// Ping 返回 *StatusCmd（恒非 nil），必须取 .Err() 判错——旧写法把 StatusCmd 当
	// error 判非空恒真，Redis 正常时 Err() 为 nil，再 .Error() 直接空指针 panic。
	if err := redisClient.Ping(ctx).Err(); err != nil {
		return dto.ServiceStatus{
			Status:  "DOWN",
			Message: "Redis ping失败: " + err.Error(),
		}
	}

	return dto.ServiceStatus{
		Status: "UP",
	}
}
