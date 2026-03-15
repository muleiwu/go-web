package service

import (
	"cnb.cool/mliev/open/go-web/pkg/server/http_server/impl"
)

type HttpServer struct {
	httpServer *impl.HttpServer
}

func (receiver *HttpServer) Run() error {
	receiver.httpServer = impl.NewHttpServer()
	receiver.httpServer.RunHttp()
	return nil
}

func (receiver *HttpServer) Stop() error {
	if receiver.httpServer == nil {
		return nil
	}
	return receiver.httpServer.Stop()
}
