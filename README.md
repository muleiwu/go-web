# Go Web 框架

[![Go Version](https://img.shields.io/badge/Go-1.25.0-blue.svg)](https://golang.org)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.10.1-green.svg)](https://github.com/gin-gonic/gin)
[![GORM](https://img.shields.io/badge/Gorm-v1.30.3-orange.svg)](https://gorm.io)

一个基于 Gin 框架的企业级 Go Web 应用模板，采用清晰的分层架构设计，内置依赖注入、健康检查、配置管理等企业级功能。

## ✨ 特性

- 🏗️ **清晰的分层架构**：Controller → Service → DAO → Model
- 🔧 **依赖注入**：基于 Assembly 模式的依赖管理
- 🌐 **多数据库支持**：MySQL、PostgreSQL
- 📦 **Redis 缓存**：内置 Redis 支持
- 📊 **健康检查**：完整的健康检查机制
- 🔧 **配置管理**：基于 Viper 的配置管理
- 📝 **结构化日志**：基于 Zap 的高性能日志
- 🚀 **优雅启停**：支持优雅关闭
- 🔄 **自动迁移**：数据库表结构自动迁移

##  📞 加群获取帮助

|                                     QQ                                      |                                 企业微信                                       |
|:---------------------------------------------------------------------------:|:--------------------------------------------------------------------------:|
| ![wechat_qr_code.png](https://static.1ms.run/dwz/image/httpsn3.inklmKc.png) | ![wechat_qr_code.png](https://static.1ms.run/dwz/image/wechat_qr_code.png) |
|        QQ群：1021660914 <br /> [点击链接加入群聊【木雷坞开源家】](https://n3.ink/lmKc)        |                                扫描上方二维码加入微信群                                |


## 🚀 快速开始

### 1. 克隆项目

```bash
git clone https://cnb.cool/mliev/open/go-web
cd go-web
```

### 2. 初始化项目

项目Fork后执行初始化脚本，自动替换模块路径：

```bash
./init.sh
```

### 3. 配置环境

```bash
# 复制配置文件
cp config.yaml.example config.yaml

# 编辑配置文件
vim config.yaml
```

### 4. 安装依赖

```bash
go mod tidy
```

### 5. 启动项目

```bash
# 开发模式
go run main.go

# 或使用 Makefile
make run
```

## 📁 项目结构

```
go-web/
├── main.go                    # 程序入口
├── go.mod                     # Go模块定义
├── go.sum                     # 依赖版本锁定
├── README.md                  # 项目说明文档
├── config.yaml.example        # 配置文件示例
├── config.yaml                # 配置文件
├── init.sh                    # 项目初始化脚本
├── Dockerfile                 # Docker构建文件
├── docker-compose.yml         # Docker编排文件
├── Makefile                   # 构建脚本
├── LICENSE                    # 许可证文件
├── app/                       # 应用核心代码
│   ├── controller/           # 控制器层
│   │   ├── base_response.go  # 统一响应封装
│   │   ├── health_controller.go # 健康检查
│   │   └── index_controller.go  # 首页控制器
│   ├── service/              # 服务层（业务逻辑）
│   ├── dao/                  # 数据访问层
│   │   └── test_demo_dao.go
│   ├── model/                # 数据模型
│   │   └── test_demo.go
│   ├── dto/                  # 数据传输对象
│   │   ├── health_dto.go
│   │   └── response_dto.go
│   └── middleware/           # 中间件
│       └── cors_middleware.go
├── cmd/                       # 命令行入口
│   └── run.go                # 启动逻辑
├── config/                   # 配置管理
│   ├── assembly.go           # 依赖注入配置
│   ├── config.go             # 基础配置
│   ├── migration.go          # 迁移配置
│   ├── server.go             # 服务器配置
│   └── autoload/             # 自动加载配置
│       ├── base.go
│       ├── database.go
│       ├── middleware.go
│       ├── redis.go
│       └── router.go
├── constants/                # 常量定义
│   └── errors.go            # 错误码定义
├── docs/                     # 文档
│   ├── PROJECT_SPECIFICATION.md
│   └── TEMPLATE_INIT.md
├── internal/                 # 内部包
│   ├── helper/               # 内部助手
│   │   └── helper.go
│   ├── interfaces/           # 接口定义
│   │   ├── assembly.go
│   │   ├── config.go
│   │   ├── database.go
│   │   ├── env.go
│   │   ├── helper.go
│   │   ├── logger.go
│   │   ├── redis.go
│   │   └── server.go
│   ├── pkg/                  # 内部包实现
│   │   ├── config/           # 配置包
│   │   ├── database/         # 数据库包
│   │   ├── demo/             # 示例包
│   │   ├── env/              # 环境变量包
│   │   ├── http_server/      # HTTP服务器包
│   │   ├── logger/           # 日志包
│   │   └── redis/            # Redis包
│   └── service/              # 内部服务
│       └── migration/        # 迁移服务
│           └── migration.go
└── util/                     # 工具函数
    ├── base_62.go
    └── generate_utils.go
```

## 🛠️ 技术栈

| 技术 | 版本 | 描述 |
|------|------|------|
| **Go** | 1.23.2 | 编程语言 |
| **Gin** | 1.10.1 | Web框架 |
| **GORM** | 1.25.12 | ORM框架 |
| **MySQL/PostgreSQL** | - | 关系型数据库 |
| **Redis** | - | 缓存数据库 |
| **Viper** | 1.19.0 | 配置管理 |
| **Zap** | 1.27.0 | 结构化日志 |

## 🏗️ 架构设计

### 分层架构
```
┌─────────────────────────────────┐
│          HTTP Layer             │
│     (Gin Router & Middleware)   │
├─────────────────────────────────┤
│        Controller Layer         │
│    (Request/Response Handling)  │
├─────────────────────────────────┤
│         Service Layer           │
│      (Business Logic)           │
├─────────────────────────────────┤
│          DAO Layer              │
│     (Data Access Objects)       │
├─────────────────────────────────┤
│         Model Layer             │
│    (Database Models)            │
└─────────────────────────────────┘
```

### 依赖注入
项目采用 Assembly 模式实现依赖注入：

```go
// Assembly 接口
type AssemblyInterface interface {
    Assembly()
}

// 配置装配顺序
func (receiver AssemblyConfig) Get() []interfaces.AssemblyInterface {
    return []interfaces.AssemblyInterface{
        assembly.Env{},      // 环境配置
        assembly.Logger{},   // 日志系统
        assembly.Database{}, // 数据库连接
        assembly.Redis{},    // Redis连接
    }
}
```

## 📋 API接口

### 健康检查
```http
# 完整健康检查
GET /health

# 简单健康检查  
GET /health/simple
```

**响应示例：**
```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "status": "UP",
    "timestamp": 1703123456,
    "services": {
      "database": {
        "status": "UP"
      },
      "redis": {
        "status": "UP"
      }
    }
  }
}
```

### 首页
```http
GET /
```

## ⚙️ 配置说明

### 配置文件示例 (config.yaml)
```yaml
# 服务配置
addr: ":8080"
mode: "debug"  # debug, release

# 数据库配置
db:
  driver: "postgresql"    # postgresql, mysql
  host: "127.0.0.1"
  port: 5432
  username: "test"
  password: "123456"
  dbname: "test"

# Redis配置
redis:
  host: "127.0.0.1"
  port: 6379
  password: ""
  db: 0

# 数据库迁移配置
database:
  halt_on_migration_failure: true
```

## 🔧 开发指南

### 编码规范
- 使用 `gofmt` 格式化代码
- 遵循 Go 官方命名规范
- 导出函数和类型必须添加注释
- 优先使用小接口，遵循单一职责原则

### 错误处理
项目使用统一的错误码和响应格式：

```go
// 错误码定义
const (
    ErrCodeSuccess      = 0   // 成功
    ErrCodeBadRequest   = 400 // 请求参数错误  
    ErrCodeUnauthorized = 401 // 未授权
    ErrCodeNotFound     = 404 // 资源不存在
    ErrCodeInternal     = 500 // 内部服务器错误
)

// 统一响应格式
type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}
```

### 数据库操作
使用 DAO 层封装数据库操作：

```go
// app/dao/user_dao.go
func GetUserByUsername(username string) (*model.User, error) {
    var user model.User
    if err := helper.Database().Where("username = ?", username).First(&user); err != nil {
        return nil, err
    }
    return &user, nil
}
```

## 🚀 部署

### Docker 部署
```bash
# 构建镜像
docker build -t go-web-app .

# 运行容器
docker run -d -p 8080:8080 go-web-app
```

### Docker Compose
```bash
# 启动服务（包含数据库和Redis）
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down
```

### 手动部署
```bash
# 编译
go build -o bin/go-web main.go

# 运行
./bin/go-web
```

## 🔨 开发命令

项目支持 Makefile 快速操作：

```bash
# 启动开发服务器
make run

# 构建项目
make build

# 运行测试
make test

# 格式化代码
make fmt

# 静态检查
make vet

# 清理构建文件
make clean
```

## 📊 监控和日志

### 健康检查端点
- **完整检查**：`GET /health` - 检查数据库、Redis等所有依赖服务
- **简单检查**：`GET /health/simple` - 仅检查服务是否启动

### 日志配置
项目使用 Zap 进行结构化日志记录：

```go
// 记录结构化日志
helper.Logger().Info("用户创建成功", 
    zap.String("username", username),
    zap.Int("userID", userID),
)
```

## 🌍 环境变量

支持通过环境变量覆盖配置文件：

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USERNAME=myuser
export DB_PASSWORD=mypass
export DB_DBNAME=mydb
export REDIS_HOST=localhost
export REDIS_PORT=6379
```

## ❓ 常见问题

### Q: 如何添加新的API路由？
A: 在 `config/autoload/router.go` 中添加路由配置，在 `app/controller/` 中实现控制器。

### Q: 如何添加新的中间件？
A: 在 `app/middleware/` 中实现中间件，然后在 `config/autoload/middleware.go` 中注册。

### Q: 数据库迁移如何工作？
A: 项目启动时会自动执行 `AutoMigrate()`，根据 Model 定义创建或更新表结构。

### Q: 如何更换数据库驱动？
A: 修改 `config.yaml` 中的 `db.driver` 配置，支持 `postgresql` 和 `mysql`。


## 📝 开发规范

详细的开发规范请参考：
- [项目规范文档](docs/PROJECT_SPECIFICATION.md)
- [模板初始化说明](docs/TEMPLATE_INIT.md)

## 🤝 贡献指南

1. Fork 项目
2. 创建功能分支：`git checkout -b feature/AmazingFeature`
3. 提交更改：`git commit -m 'Add some AmazingFeature'`
4. 推送到分支：`git push origin feature/AmazingFeature`
5. 开启 Pull Request

## 📄 许可证

本项目基于 MIT 许可证，详情请查看 [LICENSE](LICENSE) 文件。