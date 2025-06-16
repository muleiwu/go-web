# 项目开发规范文档

## 1. 项目概述

本项目是一个基于 Gin 框架的 Go Web 应用，采用分层架构设计。主要用于提供短链接服务（dwz-server）。

## 2. 技术栈

- **语言**: Go 1.23.2
- **Web框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL/PostgreSQL
- **缓存**: Redis
- **配置管理**: Viper
- **日志**: Zap
- **HTTP客户端**: go-resty

## 3. 项目目录结构规范

```
dwz-server/
├── main.go                 # 程序入口
├── go.mod                  # Go模块定义
├── go.sum                  # 依赖版本锁定
├── README.md               # 项目说明文档
├── .gitignore              # Git忽略文件
├── .gitlab-ci.yml          # CI/CD配置
├── .goreleaser.yaml        # 发布配置
├── config.yaml.example     # 配置文件示例
├── cmd/                    # 命令行入口
│   └── run.go
├── app/                    # 应用核心代码
│   ├── controller/         # 控制器层（处理HTTP请求）
│   ├── service/            # 服务层（业务逻辑）
│   ├── dao/                # 数据访问层
│   ├── model/              # 数据模型
│   ├── dto/                # 数据传输对象
│   └── middleware/         # 中间件
├── router/                 # 路由配置
├── config/                 # 配置管理
├── support/                # 基础支持服务
│   ├── db/                 # 数据库连接
│   └── logger/             # 日志配置
├── util/                   # 工具函数
└── web/                    # 静态资源
```

## 4. 命名规范

### 4.1 包名规范
- 全部小写，无下划线
- 简洁明了，避免缩写
- 目录名与包名保持一致

### 4.2 文件名规范
- 使用 CamelCase（首字母大写）
- 文件名应该清晰表达文件用途
- Controller 文件以 Controller 结尾
- Service 文件以 Service 结尾
- Model 文件直接使用实体名

### 4.3 变量和函数命名
- 公共变量/函数：CamelCase（首字母大写）
- 私有变量/函数：camelCase（首字母小写）
- 常量：UPPER_SNAKE_CASE
- 结构体：CamelCase

## 5. 代码组织规范

### 5.1 分层架构
```
HTTP请求 -> Controller -> Service -> DAO -> Database
          ↓
        DTO/Model
```

### 5.2 各层职责
- **Controller**: 处理HTTP请求，参数验证，调用Service
- **Service**: 业务逻辑处理，事务管理
- **DAO**: 数据访问，数据库操作
- **Model**: 数据模型定义
- **DTO**: 数据传输对象，API输入输出

### 5.3 依赖注入
- 使用接口定义服务契约
- 通过依赖注入减少层间耦合

## 6. 错误处理规范

### 6.1 错误码定义
```go
const (
    ErrCodeSuccess     = 0      // 成功
    ErrCodeBadRequest  = 400    // 请求参数错误
    ErrCodeUnauthorized = 401   // 未授权
    ErrCodeForbidden   = 403    // 禁止访问
    ErrCodeNotFound    = 404    // 资源不存在
    ErrCodeInternal    = 500    // 内部服务器错误
)
```

### 6.2 统一响应格式
```go
type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}
```

## 7. 数据库规范

### 7.1 表命名
- 使用 snake_case
- 表名使用复数形式
- 添加统一前缀（如项目简称）

### 7.2 字段命名
- 使用 snake_case
- 主键统一使用 id
- 时间字段：created_at, updated_at, deleted_at

### 7.3 Model定义
```go
type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Username  string    `gorm:"uniqueIndex;size:50" json:"username"`
    Password  string    `gorm:"size:100" json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

## 8. API设计规范

### 8.1 RESTful API
- GET: 获取资源
- POST: 创建资源
- PUT: 更新资源（全量）
- PATCH: 更新资源（部分）
- DELETE: 删除资源

### 8.2 URL规范
```
/api/v1/users          # 用户列表
/api/v1/users/{id}     # 特定用户
/api/v1/users/{id}/posts # 用户的文章
```

### 8.3 状态码使用
- 200: 请求成功
- 201: 创建成功
- 204: 删除成功
- 400: 请求参数错误
- 401: 未授权
- 403: 禁止访问
- 404: 资源不存在
- 500: 服务器内部错误

## 9. 日志规范

### 9.1 日志级别
- ERROR: 错误信息，需要立即处理
- WARN: 警告信息，可能存在问题
- INFO: 一般信息，记录关键业务流程
- DEBUG: 调试信息，仅在开发环境使用

### 9.2 日志格式
```go
logger.Info("用户登录成功",
    zap.String("username", username),
    zap.String("ip", clientIP),
    zap.Duration("duration", duration),
)
```

## 10. 配置管理规范

### 10.1 配置文件结构
```yaml
server:
  port: 8080
  mode: release

database:
  host: localhost
  port: 3306
  username: root
  password: password
  database: dwz

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0
```

### 10.2 环境变量
- 敏感信息使用环境变量
- 支持配置文件值被环境变量覆盖

## 11. 测试规范

### 11.1 测试分类
- 单元测试：测试单个函数/方法
- 集成测试：测试多个组件协作
- API测试：测试HTTP接口

### 11.2 测试文件命名
- 测试文件以 `_test.go` 结尾
- 测试函数以 `Test` 开头

### 11.3 覆盖率要求
- 核心业务逻辑覆盖率 >= 80%
- 工具函数覆盖率 >= 90%

## 12. 版本控制规范

### 12.1 分支管理
- main: 主分支，生产环境代码
- develop: 开发分支
- feature/*: 功能分支
- hotfix/*: 紧急修复分支

### 12.2 提交信息规范
```
type(scope): subject

body

footer
```

类型：
- feat: 新功能
- fix: 修复bug
- docs: 文档更新
- style: 代码格式调整
- refactor: 重构
- test: 测试相关
- chore: 构建过程或辅助工具的变动

## 13. 部署规范

### 13.1 构建
```bash
# 本地构建
go build -o dwz-server main.go

# 跨平台构建
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dwz-server main.go
```

### 13.2 Docker化
```dockerfile
FROM golang:1.23.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o dwz-server main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/dwz-server .
CMD ["./dwz-server"]
```

## 14. 性能规范

### 14.1 数据库
- 合理使用索引
- 避免 N+1 查询
- 使用连接池

### 14.2 缓存
- 热点数据使用 Redis 缓存
- 设置合理的过期时间
- 缓存更新策略

### 14.3 并发
- 使用 goroutine 处理并发
- 注意资源竞争和死锁
- 合理设置并发数限制

## 15. 安全规范

### 15.1 认证授权
- 使用 JWT 进行用户认证
- 实现 RBAC 权限控制
- API 限流和防刷

### 15.2 数据安全
- 敏感数据加密存储
- 密码使用哈希加盐
- SQL 注入防护

### 15.3 通信安全
- 生产环境强制 HTTPS
- 敏感接口使用签名验证
- 跨域请求控制

## 16. 监控和运维

### 16.1 健康检查
- 实现 /health 接口
- 数据库连接检查
- 外部依赖检查

### 16.2 指标监控
- 接口响应时间
- 错误率统计
- 资源使用情况

### 16.3 日志收集
- 结构化日志输出
- 日志聚合和分析
- 错误告警机制 