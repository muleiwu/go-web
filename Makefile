# 项目配置
PROJECT_NAME=dwz-server
BINARY_NAME=dwz-server
VERSION=$(shell git describe --tags --always --dirty)
BUILD_TIME=$(shell date +%Y-%m-%d\ %H:%M:%S)
GO_VERSION=$(shell go version | awk '{print $$3}')
GIT_COMMIT=$(shell git rev-parse HEAD)

# 构建标志
LDFLAGS=-ldflags "-s -w -X main.Version=${VERSION} -X main.BuildTime='${BUILD_TIME}' -X main.GoVersion='${GO_VERSION}' -X main.GitCommit='${GIT_COMMIT}'"

# 默认目标
.PHONY: help
help: ## 显示帮助信息
	@echo "可用命令："
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build
build: ## 构建应用
	@echo "构建 ${PROJECT_NAME}..."
	@go build ${LDFLAGS} -o ${BINARY_NAME} main.go
	@echo "构建完成: ${BINARY_NAME}"

.PHONY: build-linux
build-linux: ## 构建Linux版本
	@echo "构建 Linux 版本..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY_NAME}-linux main.go
	@echo "构建完成: ${BINARY_NAME}-linux"

.PHONY: build-windows
build-windows: ## 构建Windows版本
	@echo "构建 Windows 版本..."
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY_NAME}.exe main.go
	@echo "构建完成: ${BINARY_NAME}.exe"

.PHONY: build-darwin
build-darwin: ## 构建macOS版本
	@echo "构建 macOS 版本..."
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY_NAME}-darwin main.go
	@echo "构建完成: ${BINARY_NAME}-darwin"

.PHONY: build-all
build-all: build-linux build-windows build-darwin ## 构建所有平台版本

.PHONY: run
run: ## 运行应用
	@echo "启动 ${PROJECT_NAME}..."
	@go run main.go

.PHONY: dev
dev: ## 开发模式运行（需要安装air）
	@echo "启动开发模式..."
	@air

.PHONY: test
test: ## 运行测试
	@echo "运行测试..."
	@go test -v ./...

.PHONY: test-coverage
test-coverage: ## 运行测试并生成覆盖率报告
	@echo "生成测试覆盖率报告..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

.PHONY: lint
lint: ## 代码检查
	@echo "运行代码检查..."
	@go fmt ./...
	@go vet ./...
	@golangci-lint run

.PHONY: clean
clean: ## 清理构建文件
	@echo "清理构建文件..."
	@rm -f ${BINARY_NAME}*
	@rm -f coverage.out coverage.html
	@rm -rf dist/
	@echo "清理完成"

.PHONY: deps
deps: ## 安装依赖
	@echo "安装依赖..."
	@go mod download
	@go mod tidy

.PHONY: update-deps
update-deps: ## 更新依赖
	@echo "更新依赖..."
	@go get -u ./...
	@go mod tidy

.PHONY: docker-build
docker-build: ## 构建Docker镜像
	@echo "构建 Docker 镜像..."
	@docker build -t ${PROJECT_NAME}:${VERSION} .
	@docker tag ${PROJECT_NAME}:${VERSION} ${PROJECT_NAME}:latest
	@echo "Docker 镜像构建完成"

.PHONY: docker-run
docker-run: ## 运行Docker容器
	@echo "运行 Docker 容器..."
	@docker run --rm -p 8080:8080 --name ${PROJECT_NAME} ${PROJECT_NAME}:latest

.PHONY: docker-compose-up
docker-compose-up: ## 启动docker-compose服务
	@echo "启动 docker-compose 服务..."
	@docker-compose up -d

.PHONY: docker-compose-down
docker-compose-down: ## 停止docker-compose服务
	@echo "停止 docker-compose 服务..."
	@docker-compose down

.PHONY: docker-compose-logs
docker-compose-logs: ## 查看docker-compose日志
	@docker-compose logs -f

.PHONY: release
release: ## 发布版本（使用goreleaser）
	@echo "发布版本..."
	@goreleaser release --clean

.PHONY: release-snapshot
release-snapshot: ## 快照发布（不推送到仓库）
	@echo "快照发布..."
	@goreleaser release --snapshot --clean

.PHONY: install-tools
install-tools: ## 安装开发工具
	@echo "安装开发工具..."
	@go install github.com/air-verse/air@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/goreleaser/goreleaser/v2@latest
	@echo "开发工具安装完成"

.PHONY: generate
generate: ## 生成代码
	@echo "生成代码..."
	@go generate ./...

.PHONY: migrate
migrate: ## 运行数据库迁移
	@echo "运行数据库迁移..."
	@go run main.go migrate

.PHONY: config
config: ## 复制配置文件示例
	@echo "复制配置文件..."
	@cp config.yaml.example config.yaml
	@echo "请编辑 config.yaml 文件配置您的环境"

.PHONY: check
check: lint test ## 检查代码质量（格式化、静态检查、测试）

.PHONY: all
all: clean deps check build ## 完整构建流程 