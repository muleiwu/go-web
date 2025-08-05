package config

import (
	"cnb.cool/mliev/examples/go-web/internal/assembly"
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
)

type AssemblyConfig struct {
}

func (receiver AssemblyConfig) Get() []interfaces.AssemblyInterface {
	return []interfaces.AssemblyInterface{
		assembly.Env{},
		assembly.Logger{},
	}
}
