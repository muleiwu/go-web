# 路由使用指南

## 概述

go-web 框架在 Gin 之上提供了一层路由抽象，通过 `RouterContextInterface` 屏蔽底层框架细节。所有路由 handler 统一使用 `HandlerFunc` 签名，不直接依赖 `*gin.Context`。

路由系统支持三种路由模式：
- **标准路由** — 基于 Gin 的精确路径匹配（含路径参数 `:id`）
- **路由分组** — 共享前缀和中间件的路由组
- **正则路由** — 基于正则表达式的高级路由匹配，支持命名捕获组

## 快速开始

路由在 `config/autoload/router.go` 中通过 `InitConfig` 注册：

```go
package autoload

import (
    httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
    "cnb.cool/mliev/open/go-web/app/controller"
)

type Router struct{}

func (receiver Router) InitConfig() map[string]any {
    return map[string]any{
        "http.router": func(router httpInterfaces.RouterInterface) {
            router.GET("/", controller.IndexController{}.GetIndex)
            router.GET("/health", controller.HealthController{}.GetHealth)
        },
    }
}
```

## 路由注册

`RouterInterface` 提供以下 HTTP 方法注册：

```go
router.GET(path string, handler HandlerFunc)
router.POST(path string, handler HandlerFunc)
router.PUT(path string, handler HandlerFunc)
router.DELETE(path string, handler HandlerFunc)
router.PATCH(path string, handler HandlerFunc)
router.HEAD(path string, handler HandlerFunc)
router.OPTIONS(path string, handler HandlerFunc)
```

路径参数使用 Gin 的 `:param` 语法：

```go
router.GET("/users/:id", controller.UserController{}.GetUser)
```

在 handler 中通过 `c.Param("id")` 获取参数值。

## 路由分组

使用 `Group` 方法创建共享前缀的路由组：

```go
"http.router": func(router httpInterfaces.RouterInterface) {
    // /api/v1/users, /api/v1/posts ...
    v1 := router.Group("/api/v1")
    {
        v1.GET("/users", controller.UserController{}.List)
        v1.POST("/users", controller.UserController{}.Create)
    }

    // 嵌套分组
    health := router.Group("/health")
    {
        health.GET("", controller.HealthController{}.GetHealth)
        health.GET("/simple", controller.HealthController{}.GetHealthSimple)
    }
}
```

分组支持独立的中间件：

```go
admin := router.Group("/admin")
admin.Use(middleware.AuthMiddleware())
{
    admin.GET("/dashboard", controller.AdminController{}.Dashboard)
}
```

## 正则路由

当 Gin 内置的路径参数无法满足需求时，使用 `RegexGroup` 创建正则路由组：

```go
"http.router": func(router httpInterfaces.RouterInterface) {
    regex := router.RegexGroup("/api")
    {
        // 匹配 /api/users/123
        regex.GET(`^/api/users/(?P<id>\d+)$`, controller.UserController{}.GetUser)

        // 匹配 /api/posts/456/comments
        regex.POST(`^/api/posts/(?P<postId>\d+)/comments$`, controller.PostController{}.AddComment)

        // 匹配所有 HTTP 方法
        regex.Any(`^/api/proxy/(?P<path>.+)$`, controller.ProxyController{}.Handle)
    }
}
```

### 正则路由特性

- **命名捕获组**：`(?P<name>pattern)` 中的 `name` 会自动写入路径参数，通过 `c.Param("name")` 获取
- **Handler 链**：正则路由支持多个 handler，依次执行：
  ```go
  regex.GET(`^/api/resource/(?P<id>\d+)$`, authHandler, resourceHandler)
  ```
- **惰性挂载**：同一前缀下的所有正则路由共享一个 Gin 通配路由，首次注册时自动挂载
- **无匹配返回 404**：请求未命中任何正则规则时返回 `404 not found`

### 可用方法

```go
regex.GET(pattern string, handlers ...HandlerFunc)
regex.POST(pattern string, handlers ...HandlerFunc)
regex.PUT(pattern string, handlers ...HandlerFunc)
regex.DELETE(pattern string, handlers ...HandlerFunc)
regex.PATCH(pattern string, handlers ...HandlerFunc)
regex.HEAD(pattern string, handlers ...HandlerFunc)
regex.OPTIONS(pattern string, handlers ...HandlerFunc)
regex.Any(pattern string, handlers ...HandlerFunc)    // 匹配所有 HTTP 方法
```

## 中间件

### 全局中间件

在 `config/autoload/middleware.go` 中注册全局中间件：

```go
package autoload

import (
    httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
    "cnb.cool/mliev/open/go-web/app/middleware"
)

type Middleware struct{}

func (receiver Middleware) InitConfig() map[string]any {
    return map[string]any{
        "http.middleware": []httpInterfaces.MiddlewareFunc{
            middleware.CorsMiddleware(),
        },
    }
}
```

### 分组中间件

通过 `Use` 方法为路由组添加中间件：

```go
api := router.Group("/api")
api.Use(middleware.AuthMiddleware())
```

### 编写自定义中间件

中间件签名为 `MiddlewareFunc`，即 `func(RouterContextInterface)`。必须调用 `c.Next()` 继续执行后续 handler，或调用 `c.Abort()` / `c.AbortWithStatus()` 中止请求。

```go
package middleware

import (
    "net/http"
    httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
)

func AuthMiddleware() httpInterfaces.MiddlewareFunc {
    return func(c httpInterfaces.RouterContextInterface) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        // 验证 token，将用户信息存入上下文
        c.Set("userId", "12345")
        c.Next()
    }
}
```

## 控制器编写

### Handler 签名

所有路由 handler 的签名为：

```go
type HandlerFunc func(RouterContextInterface)
```

### 控制器结构

推荐嵌入 `BaseResponse` 以使用统一的响应方法：

```go
package controller

import (
    "net/http"
    httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
)

type UserController struct {
    BaseResponse
}

func (receiver UserController) GetUser(c httpInterfaces.RouterContextInterface) {
    id := c.Param("id")
    // 业务逻辑...
    receiver.Success(c, map[string]any{"id": id, "name": "张三"})
}
```

### BaseResponse 响应方法

```go
// 成功响应 → {"code": 0, "message": "ok", "data": ...}
receiver.Success(c, data)

// 成功响应（自定义消息）
receiver.SuccessWithMessage(c, "操作成功", data)

// 错误响应 → {"code": 1001, "message": "参数错误"}
receiver.Error(c, constants.ErrCodeInvalidParams, "参数错误")

// 错误响应（携带数据）
receiver.ErrorWithData(c, 400, "验证失败", validationErrors)
```

错误码为 400-599 范围时，HTTP 状态码与错误码一致；否则 HTTP 状态码为 200。

## 请求上下文 API 速查表

`RouterContextInterface` 提供以下方法：

### 请求参数

| 方法 | 说明 |
|------|------|
| `Param(key string) string` | 获取路径参数（`:id` 或正则命名捕获组） |
| `Query(key string) string` | 获取 URL 查询参数 |
| `DefaultQuery(key, defaultValue string) string` | 获取查询参数，不存在时返回默认值 |
| `PostForm(key string) string` | 获取表单数据 |
| `ShouldBindJSON(obj any) error` | 解析 JSON 请求体到结构体 |

### 响应

| 方法 | 说明 |
|------|------|
| `JSON(code int, obj any)` | 返回 JSON 响应 |
| `String(code int, format string, values ...any)` | 返回文本响应 |
| `HTML(code int, name string, obj any)` | 渲染 HTML 模板 |
| `Data(code int, contentType string, data []byte)` | 返回原始数据 |
| `Redirect(code int, location string)` | HTTP 重定向 |

### HTTP 头和 Cookie

| 方法 | 说明 |
|------|------|
| `GetHeader(key string) string` | 获取请求头 |
| `SetHeader(key, value string)` | 设置响应头 |
| `Cookie(name string) (string, error)` | 获取 Cookie |
| `SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)` | 设置 Cookie |

### 请求信息

| 方法 | 说明 |
|------|------|
| `Path() string` | 请求路径 |
| `FullPath() string` | 完整路由模式 |
| `Method() string` | HTTP 方法 |
| `ClientIP() string` | 客户端 IP |

### 上下文存取

| 方法 | 说明 |
|------|------|
| `Set(key string, value any)` | 存储值到上下文 |
| `Get(key string) any` | 从上下文取值 |
| `GetString(key string) string` | 从上下文取字符串值 |

### 流程控制

| 方法 | 说明 |
|------|------|
| `Next()` | 执行下一个 handler（中间件中使用） |
| `Abort()` | 中止请求处理 |
| `AbortWithStatus(code int)` | 中止并返回指定 HTTP 状态码 |
