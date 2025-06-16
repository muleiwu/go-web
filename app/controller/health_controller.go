package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"mliev.com/template/go-web/constants"
	"mliev.com/template/go-web/helper"
	"time"
)

// HealthStatus 健康状态结构
type HealthStatus struct {
	Status    string                 `json:"status"`
	Timestamp int64                  `json:"timestamp"`
	Services  map[string]interface{} `json:"services"`
}

// ServiceStatus 服务状态
type ServiceStatus struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// GetHealth 健康检查接口
func GetHealth(c *gin.Context) {
	healthStatus := HealthStatus{
		Status:    "UP",
		Timestamp: time.Now().Unix(),
		Services:  make(map[string]interface{}),
	}

	// 检查数据库连接
	dbStatus := checkDatabase()
	healthStatus.Services["database"] = dbStatus

	// 检查Redis连接
	redisStatus := checkRedis()
	healthStatus.Services["redis"] = redisStatus

	// 如果任何服务不健康，整体状态设为DOWN
	if dbStatus.Status == "DOWN" || redisStatus.Status == "DOWN" {
		healthStatus.Status = "DOWN"
		Error(c, constants.ErrCodeUnavailable, "服务不健康")
		return
	}

	Success(c, healthStatus)
}

// GetHealthSimple 简单健康检查接口
func GetHealthSimple(c *gin.Context) {
	Success(c, gin.H{
		"status":    "UP",
		"timestamp": time.Now().Unix(),
	})
}

// checkDatabase 检查数据库连接
func checkDatabase() ServiceStatus {
	database := helper.GetDB()
	if database == nil {
		return ServiceStatus{
			Status:  "DOWN",
			Message: "数据库连接失败",
		}
	}

	sqlDB, err := database.DB()
	if err != nil {
		return ServiceStatus{
			Status:  "DOWN",
			Message: "获取数据库连接失败: " + err.Error(),
		}
	}

	if err := sqlDB.Ping(); err != nil {
		return ServiceStatus{
			Status:  "DOWN",
			Message: "数据库ping失败: " + err.Error(),
		}
	}

	return ServiceStatus{
		Status: "UP",
	}
}

// checkRedis 检查Redis连接
func checkRedis() ServiceStatus {
	redis := helper.GetRedis()
	if redis == nil {
		return ServiceStatus{
			Status:  "DOWN",
			Message: "Redis连接失败",
		}
	}

	ctx := context.Background()
	if err := redis.Ping(ctx).Err(); err != nil {
		return ServiceStatus{
			Status:  "DOWN",
			Message: "Redis ping失败: " + err.Error(),
		}
	}

	return ServiceStatus{
		Status: "UP",
	}
}
