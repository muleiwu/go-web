package config

import (
	"cnb.cool/mliev/open/go-web/config/autoload"
	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	"cnb.cool/mliev/open/go-web/pkg/server/http_server/service"
	"cnb.cool/mliev/open/go-web/pkg/server/migration"
)

type Server struct {
}

func (receiver Server) Get() []interfaces.ServerInterface {
	return []interfaces.ServerInterface{
		&migration.Migration{
			Migration: autoload.Migration{}.Get(),
		},
		&service.HttpServer{},
	}
}
