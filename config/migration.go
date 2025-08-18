package config

type Migration struct {
	// 是否启用自动迁移
	AutoMigrate bool
	// 需要迁移的模型列表
	Models []any
}

func (receiver Migration) Get() []any {
	if receiver.AutoMigrate {
		return receiver.Models
	}
	return []any{
		//&model.TestDemo{},
	}
}
