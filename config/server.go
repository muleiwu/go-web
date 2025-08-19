package config

import (
	"cnb.cool/mliev/examples/go-web/internal/helper"
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	"cnb.cool/mliev/examples/go-web/internal/pkg/http_server/service"
	"cnb.cool/mliev/examples/go-web/internal/service/migration"
)

type Server struct {
	Helper *helper.Helper
}

func (receiver Server) Get() []interfaces.ServerInterface {
	return []interfaces.ServerInterface{
		&migration.Migration{
			Helper:    receiver.Helper,
			Migration: Migration{}.Get(),
		},
		&service.HttpServer{
			Helper: receiver.Helper,
		},
	}
}
