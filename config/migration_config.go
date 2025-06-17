package config

import "cnb.cool/mliev/examples/go-web/app/model"

type MigrationConfig struct {
	// 是否启用自动迁移
	AutoMigrate bool
	// 需要迁移的模型列表
	Models []any
}

func (receiver MigrationConfig) Get() []any {
	if receiver.AutoMigrate {
		return receiver.Models
	}
	return []any{
		&model.TestDemo{},
	}
}
