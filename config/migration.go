package config

type Migration struct {
}

func (receiver Migration) Get() []any {
	return []any{
		//&model.TestDemo{},
	}
}
