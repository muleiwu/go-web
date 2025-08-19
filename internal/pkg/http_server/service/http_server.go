package service

import (
	"cnb.cool/mliev/examples/go-web/internal/helper"
	"cnb.cool/mliev/examples/go-web/internal/pkg/http_server/impl"
)

type HttpServer struct {
	Helper *helper.Helper
}

func (receiver *HttpServer) Run() error {

	newHttpServer := impl.NewHttpServer(receiver.Helper)

	newHttpServer.RunHttp()
	return nil
}
