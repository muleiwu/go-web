package autoload

import envInterface "cnb.cool/mliev/examples/go-web/internal/interfaces"

type Migration struct {
}

func (receiver Migration) Get() []any {
	return []any{
		//&model.TestDemo{},
	}
}

func (receiver Migration) InitConfig(helper envInterface.HelperInterface) map[string]any {
	return map[string]any{
		"database.migration": receiver.Get(),
	}
}
