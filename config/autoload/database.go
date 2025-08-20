package autoload

import envInterface "cnb.cool/mliev/examples/go-web/internal/interfaces"

type Database struct {
}

func (receiver Database) InitConfig(env envInterface.EnvInterface) map[string]any {
	return map[string]any{
		"database.driver":   env.GetString("database.driver", "mysql"),
		"database.host":     env.GetString("database.host", "127.0.0.1"),
		"database.port":     env.GetInt("database.port", 3306),
		"database.dbname":   env.GetString("database.dbname", "test"),
		"database.username": env.GetString("database.username", "test"),
		"database.password": env.GetString("database.password", "123456"),
	}
}
