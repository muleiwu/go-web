# 数据库迁移

本项目使用 [goose](https://github.com/pressly/goose) 进行版本化数据库迁移。

## 快速开始

```bash
# 创建新的迁移文件
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

## 迁移文件格式

执行 `create` 命令后会在本目录生成形如 `20260325120000_create_users_table.sql` 的文件，格式如下：

```sql
-- +goose Up
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS users;
```

### 规则

- `-- +goose Up` 标记向上迁移（创建/修改），`-- +goose Down` 标记回滚操作
- 每个文件必须同时包含 Up 和 Down 两部分
- Down 应能完全撤销 Up 的变更
- 文件按时间戳顺序执行，不要手动修改文件名中的时间戳

### 单语句与多语句

默认每条 SQL 独立执行。如果需要在一个事务中执行多条语句：

```sql
-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL
);

INSERT INTO orders (user_id, amount) VALUES (1, 100.00);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS orders;
```

## 自动迁移

服务启动时会自动执行 `goose.Up`，应用所有未执行的迁移。如需跳过，将数据库驱动设为 `memory`。

## 配置

在 `config.yaml` 中可自定义迁移目录：

```yaml
database:
  migration:
    dir: "migrations"
```

也可通过环境变量覆盖：`DATABASE_MIGRATION_DIR=migrations`
