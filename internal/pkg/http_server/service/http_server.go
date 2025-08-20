package service

import (
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	"cnb.cool/mliev/examples/go-web/internal/pkg/http_server/impl"
)

type HttpServer struct {
	Helper interfaces.HelperInterface
}

func (receiver *HttpServer) Run() error {

	newHttpServer := impl.NewHttpServer(receiver.Helper)

	newHttpServer.RunHttp()
	return nil
}
