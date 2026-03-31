# Go Web 框架

[![Go Version](https://img.shields.io/badge/Go-1.25.0-blue.svg)](https://golang.org)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.11.0-green.svg)](https://github.com/gin-gonic/gin)
[![GORM](https://img.shields.io/badge/Gorm-v1.31.1-orange.svg)](https://gorm.io)

一个基于 Gin 框架的企业级 Go Web 应用模板，采用清晰的分层架构设计，内置依赖注入容器、可插拔驱动系统、版本化数据库迁移等企业级功能。

## 特性

- **清晰的分层架构**：Controller → Service → DAO → Model
- **依赖注入容器**：懒加载单例、优先级覆盖、反射注入
- **可插拔驱动系统**：数据库、日志、缓存、Redis 均支持运行时切换驱动
- **多数据库支持**：MySQL、PostgreSQL、SQLite、Memory
- **版本化数据库迁移**：基于 Goose 的 SQL 迁移，支持 up/down/status/create
- **Redis 与缓存**：内置 Redis 支持，缓存驱动可选 redis/memory/none
- **健康检查**：完整的健康检查机制
- **配置管理**：基于 Viper 的配置管理，支持环境变量覆盖
- **结构化日志**：基于 Zap 的高性能日志（development/production 驱动）
- **优雅启停**：SIGINT/SIGTERM 优雅关闭，SIGHUP 热重载
- **框架无关控制器**：控制器接受 `RouterContextInterface`，不直接依赖 Gin
- **正则路由**：支持正则表达式路由与命名捕获组

## 加群获取帮助

|                                     QQ                                      |                                 企业微信                                       |
|:---------------------------------------------------------------------------:|:--------------------------------------------------------------------------:|
| ![wechat_qr_code.png](https://static.1ms.run/dwz/image/httpsn3.inklmKc.png) | ![wechat_qr_code.png](https://static.1ms.run/dwz/image/wechat_qr_code.png) |
|        QQ群：1021660914 <br /> [点击链接加入群聊【木雷坞开源家】](https://n3.ink/lmKc)        |                                扫描上方二维码加入微信群                                |

## 快速开始

### 1. 克隆项目

```bash
git clone https://cnb.cool/mliev/open/go-web
cd go-web
```

### 2. 初始化项目

Fork 后执行初始化脚本，自动替换模块路径：

```bash
./init.sh
```

### 3. 配置环境

```bash
cp config.yaml.example config.yaml
vim config.yaml
```

### 4. 安装依赖

```bash
go mod tidy
```

### 5. 启动项目

```bash
# 前台启动
go run main.go start

# 守护进程模式
go run main.go start -d

# 停止 / 重启 / 热重载 / 状态
go run main.go stop
go run main.go restart
go run main.go reload
go run main.go status
```

## 项目结构

```
go-web/
├── main.go                          # 程序入口，嵌入静态资源
├── go.mod / go.sum                  # Go 模块定义
├── config.yaml.example              # 配置文件示例
├── init.sh                          # 项目初始化脚本（替换模块路径）
├── Dockerfile                       # Docker 构建文件
├── app/                             # 应用层代码
│   ├── controller/                  # 控制器（接受 RouterContextInterface）
│   │   ├── base_response.go         # 统一响应封装（Success/Error）
│   │   ├── health_controller.go     # 健康检查
│   │   └── index_controller.go      # 首页
│   ├── service/                     # 服务层（业务逻辑）
│   ├── dao/                         # 数据访问层
│   ├── model/                       # 数据模型（GORM）
│   ├── dto/                         # 数据传输对象
│   │   ├── health_dto.go
│   │   └── response_dto.go
│   ├── constants/                   # 常量与错误码
│   │   └── errors.go
│   └── middleware/                   # 中间件
│       └── cors_middleware.go
├── cmd/                             # 命令行入口
│   ├── run.go                       # 启动逻辑（信号处理、热重载）
│   ├── options.go                   # Functional Options
│   └── migrate/                     # 数据库迁移 CLI
│       └── main.go
├── config/                          # 配置管理
│   ├── app.go                       # AppProvider 实现（装配链 + 服务链）
│   ├── config.go                    # 配置加载清单
│   └── autoload/                    # 自动加载配置
│       ├── app.go                   # 应用配置
│       ├── cache.go                 # 缓存配置
│       ├── database.go              # 数据库配置
│       ├── http.go                  # HTTP 配置
│       ├── middleware.go            # 中间件注册
│       ├── migration.go             # 迁移配置
│       ├── redis.go                 # Redis 配置
│       ├── router.go                # 路由注册
│       └── static_fs.go            # 静态文件配置
├── migrations/                      # SQL 迁移文件目录
├── pkg/                             # 框架核心
│   ├── container/                   # DI 容器（懒加载单例 + 优先级覆盖）
│   ├── contract/                    # 服务契约（类型别名）
│   ├── driver/                      # 泛型驱动管理器
│   ├── helper/                      # 全局访问器（GetDatabase/GetLogger/...）
│   ├── interfaces/                  # 核心接口（AppProvider/Assembly/Server）
│   └── server/                      # 各基础设施实现
│       ├── cache/                   # 缓存（redis/memory/none 驱动）
│       ├── config/                  # 配置服务
│       ├── database/                # 数据库（mysql/postgresql/sqlite/memory 驱动）
│       ├── env/                     # 环境变量（Viper）
│       ├── http_server/             # HTTP 服务器（Gin）
│       ├── logger/                  # 日志（development/production 驱动）
│       ├── migration/               # 数据库迁移（Goose）
│       ├── redis/                   # Redis 客户端
│       └── reload/                  # 热重载信号通道
├── static/                          # 静态资源（go:embed）
├── templates/                       # 模板文件（go:embed）
└── docs/                            # 文档
```

## 技术栈

| 技术 | 版本 | 用途 |
|------|------|------|
| **Go** | 1.25.0 | 编程语言 |
| **Gin** | 1.11.0 | Web 框架 |
| **GORM** | 1.31.1 | ORM |
| **Goose** | 3.27.0 | 数据库迁移 |
| **Viper** | 1.21.0 | 配置管理 |
| **Zap** | 1.27.1 | 结构化日志 |
| **go-redis** | 9.17.3 | Redis 客户端 |
| **Cobra** | 1.10.2 | CLI 框架 |
| **gomander** | 1.0.0 | 进程管理（start/stop/reload） |

## 架构设计

### 启动流程

```
main.go → gomander.Run() → cmd.Start(WithTemplateFs, WithWebStaticFs, WithApp)
  → Assembly 链: Env → Config → Logger → Database → Redis → Cache
  → Server 链:   Migration(goose.Up) → HttpServer(Gin)
  → 信号监听:     SIGINT/SIGTERM → 优雅关闭 | SIGHUP → 热重载
```

### 分层架构

```
┌─────────────────────────────────┐
│        HTTP Layer (Gin)         │
├─────────────────────────────────┤
│      Controller Layer           │
│  (RouterContextInterface)       │
├─────────────────────────────────┤
│       Service Layer             │
│     (Business Logic)            │
├─────────────────────────────────┤
│        DAO Layer                │
│    (Data Access Objects)        │
├─────────────────────────────────┤
│       Model Layer               │
│    (GORM Models)                │
└─────────────────────────────────┘
```

### DI 容器

项目使用自定义 DI 容器（`pkg/container/`），支持：

- **懒加载单例**：首次 `Get` 时创建实例
- **优先级覆盖**：同名 Provider 高优先级覆盖低优先级
- **反射注入**：通过 `inject:""` 标签按字段类型自动注入（`inject:"-"` 跳过）
- **生命周期钩子**：`Initializable`（初始化）/ `Destroyable`（销毁）
- **热重载支持**：`ResetAll()` 重置所有实例后重新装配

```go
// 获取服务
db := container.MustGet[*gorm.DB]()

// 或使用 helper 快捷方式
db := helper.GetDatabase()
logger := helper.GetLogger()
cache := helper.GetCache()
```

### 可插拔驱动

基于泛型 `driver.Manager[T]`，每种基础设施支持多个驱动实现，在配置文件中通过名称切换：

| 服务 | 可用驱动 |
|------|---------|
| 数据库 | `mysql`、`postgresql`、`sqlite`、`memory` |
| 日志 | `development`、`production` |
| 缓存 | `redis`、`memory`、`none` |

## 数据库迁移

项目使用 [Goose](https://github.com/pressly/goose) 进行版本化 SQL 迁移。

```bash
# 创建迁移文件
go run cmd/migrate/main.go create create_users_table

# 执行所有未应用的迁移
go run cmd/migrate/main.go up

# 回滚最近一次迁移
go run cmd/migrate/main.go down

# 查看迁移状态
go run cmd/migrate/main.go status

# 回滚并重新执行最近一次迁移
go run cmd/migrate/main.go redo
```

迁移文件格式：

```sql
-- +goose Up
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS users;
```

服务启动时会自动执行 `goose.Up`，应用所有未执行的迁移。详见 `migrations/README.md`。

## API 接口

### 健康检查

```http
# 完整健康检查（数据库 + Redis）
GET /health

# 简单健康检查
GET /health/simple
```

**响应格式：**

```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "status": "UP",
    "timestamp": 1703123456,
    "services": {
      "database": { "status": "UP" },
      "redis": { "status": "UP" }
    }
  }
}
```

### 统一响应格式

所有接口使用统一响应结构：

```json
{
  "code": 0,
  "message": "操作成功",
  "data": {}
}
```

错误码定义在 `app/constants/errors.go`。控制器通过嵌入 `BaseResponse` 使用 `Success()` / `Error()` 方法。

## 配置说明

### 配置文件 (config.yaml)

```yaml
# 数据库配置
database:
  driver: "mysql"           # mysql, postgresql, sqlite, memory
  host: "127.0.0.1"
  port: 3306
  username: "root"
  password: "password"
  dbname: "mydb"
  migration:
    dir: "migrations"       # SQL 迁移文件目录

# Redis 配置
redis:
  host: "127.0.0.1"
  port: 6379
  password: ""
  db: 0

# 应用配置
app:
  mode: "debug"             # debug, release
```

### 环境变量

所有配置项均可通过环境变量覆盖，使用 `_` 替代 `.` 作为分隔符：

```bash
export DATABASE_DRIVER=postgresql
export DATABASE_HOST=localhost
export DATABASE_PORT=5432
export DATABASE_USERNAME=myuser
export DATABASE_PASSWORD=mypass
export DATABASE_DBNAME=mydb
export DATABASE_MIGRATION_DIR=migrations
export REDIS_HOST=localhost
export REDIS_PORT=6379
```

## 部署

### Docker

```bash
# 构建镜像
docker build -t go-web-app .

# 运行容器
docker run -d -p 8080:8080 go-web-app
```

### 手动部署

```bash
# 编译
go build -o bin/go-web main.go

# 前台运行
./bin/go-web start

# 守护进程运行
./bin/go-web start -d
```

## 开发指南

### 添加新路由

在 `config/autoload/router.go` 中注册路由，在 `app/controller/` 中实现控制器：

```go
// config/autoload/router.go
router.GET("/users", controller.UserController{}.List)
router.POST("/users", controller.UserController{}.Create)
```

### 添加新中间件

在 `app/middleware/` 中实现，在 `config/autoload/middleware.go` 中注册：

```go
// config/autoload/middleware.go
"http.middleware": []httpInterfaces.MiddlewareFunc{
    middleware.CorsMiddleware(),
    middleware.AuthMiddleware(),
},
```

### 编码规范

- 使用 `gofmt` 格式化代码
- 导入顺序：标准库、第三方库、项目包（空行分隔）
- 文件名使用 snake_case
- 控制器方法接受 `RouterContextInterface`，不直接使用 `*gin.Context`

## 贡献指南

1. Fork 项目
2. 创建功能分支：`git checkout -b feature/AmazingFeature`
3. 提交更改：`git commit -m 'Add some AmazingFeature'`
4. 推送到分支：`git push origin feature/AmazingFeature`
5. 开启 Pull Request

## 许可证

本项目基于 MIT 许可证，详情请查看 [LICENSE](LICENSE) 文件。
