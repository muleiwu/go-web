# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Enterprise-grade Go Web framework built on Gin, with a custom DI container, pluggable driver system, and clean layered architecture. Module: `cnb.cool/mliev/open/go-web`. Requires Go 1.25.0.

## Build & Run Commands

```bash
go mod tidy          # Install dependencies
go run main.go       # Run in development
go build -o bin/go-web main.go  # Build binary
go fmt ./...         # Format code
go vet ./...         # Vet
go test ./...        # Run all tests
go test -run TestFoo ./path/to/package  # Run a single test

# Docker
docker build -t go-web-app .
docker run -d -p 8080:8080 go-web-app
```

Configuration: copy `config.yaml.example` to `config.yaml` before running. No Makefile exists (README references are outdated).

## Architecture

### Startup Flow
```
main.go → gomander.Run() → cmd.Start(functional options) → Assembly chain → Servers
```

1. `main.go` embeds `templates/` and `static/` via `//go:embed`, calls `cmd.Start()` with `WithTemplateFs`, `WithWebStaticFs`, `WithApp`
2. `cmd/run.go` sets up signal handling, runs the Assembly chain, then starts Servers
3. Assembly order: Env → Config → Logger → Database → Redis → Cache
4. Servers: Migration (auto-migrate DB tables), then HttpServer
5. Graceful shutdown on SIGINT/SIGTERM; hot reload (full container reset + reassembly) on SIGHUP or via `reload.GetReloadChan()`

### DI Container (`pkg/container/`)

Global singleton container with lazy-loading. Services are registered as `Provider` implementations (with `Type() reflect.Type`, `Build() any`, `Priority() int`). Higher-priority providers override lower ones. Keys are `reflect.Type`, not strings.

- `container.Register(provider)` — register a provider (typically in assembly)
- `container.Get[T]()` / `container.MustGet[T]()` — retrieve by type parameter (no name needed)
- `container.Inject(target)` — auto-inject struct fields tagged `inject:""` (by field type, not name; `inject:"-"` to skip)
- `SimpleProvider` — wraps an existing instance; `LazyProvider` — defers creation via factory func
- `Initializable` / `Destroyable` interfaces for lifecycle hooks
- `DependencyAware` interface — providers can declare `DependsOn() []reflect.Type` for topological sorting
- `container.ResetAll()` — used during SIGHUP reload to force re-creation
- Runtime circular dependency detection via goroutine ID tracking

### Driver Manager (`pkg/driver/`)

Generic `Manager[T]` for runtime-selectable implementations. Each domain (database, logger, cache, redis, static files) has its own manager under `pkg/server/*/driver/manager.go`. Drivers are registered via `Extend(name, factory)` and created via `Make(name, config)`.

Available drivers:
- Database: mysql, postgresql, sqlite, memory
- Logger: development, production
- Cache: redis, memory, none
- Redis: redis

### Application Layers (`app/`)

Controller → Service → DAO → Model, with DTOs for API I/O.

- Controllers accept `RouterContextInterface` (not `*gin.Context`) — framework-agnostic
- Embed `BaseResponse` struct for unified `Success()`/`Error()` response methods
- Response format: `{"code": 0, "message": "...", "data": ...}`
- Error codes in `app/constants/errors.go`

### Configuration System

Two-layer config: **Env** (Viper, reads `config.yaml` + env vars with `.` → `_` replacement) and **Config** (in-memory `gsr.Provider` populated from `InitConfig` declarations).

- `config/config.go` — lists all `InitConfig` implementations to load
- `config/autoload/*.go` — each file implements `InitConfig` returning `map[string]any` of default config values read from env
- Config keys use dot notation: `app.mode`, `database.driver`, `http.addr`, `redis.host`, etc.

### AppProvider Pattern (`config/app.go`)

`App` struct implements `AppProvider` interface with two methods:
- `Assemblies()` — returns ordered assembly chain
- `Servers()` — returns servers to run (migration + HTTP)

Custom apps can provide their own `AppProvider` via `cmd.WithApp()`.

### Routing & Middleware

- Routes: `config/autoload/router.go` — register routes via `RouterInterface` callback in `http.router` config key
- Middleware: `config/autoload/middleware.go` — `[]MiddlewareFunc` in `http.middleware` config key
- Standard routes: `router.GET/POST/PUT/DELETE/PATCH/HEAD/OPTIONS`
- Route groups: `router.Group(path)`
- Regex routes: `router.RegexGroup(prefix)` with named capture groups

### Global Helpers (`pkg/helper/`)

Convenience accessors that call `container.MustGet[T]()`:
- `helper.GetDatabase()` → `*gorm.DB`
- `helper.GetLogger()` → logger
- `helper.GetRedis()` → redis client
- `helper.GetConfig()` → config provider
- `helper.GetCache()` → cache
- `helper.GetEnv()` → env reader

## Code Conventions

- Import order: stdlib, third-party, project packages (separated by blank lines)
- File names: snake_case
- Chinese language in commit messages and user-facing strings
- `gomander` wraps cobra for CLI; entry point is `gomander.Run(func() { cmd.Start(...) })`
- Key third-party libs: `github.com/muleiwu/gsr` (service registry interfaces), `github.com/muleiwu/anyto` (type conversion), `github.com/muleiwu/go-cache` (caching), `github.com/muleiwu/golog` (logging)
- Existing tests are in `pkg/container/` (container, assembly, topo sort); no application-layer tests yet
- Health endpoints: `/health` (checks DB + Redis), `/health/simple` (just up/down)