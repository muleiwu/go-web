package autoload

type Migration struct {
}

func (receiver Migration) Get() []any {
	return []any{
		//&model.TestDemo{},
	}
}

func (receiver Migration) InitConfig() map[string]any {
	return map[string]any{
		"database.migration": receiver.Get(),
	}
}
