package app

import "github.com/blockfint/di-example-go/app/server"

type Application struct {
	Server *server.Server
}

func NewApplication(server *server.Server) *Application {
	return &Application{server}
}
