package sample

import (
	"net"
	"github.com/acidlemon/rocket"
)

type WebApp struct {
	rocket.WebApp
}

func NewWebApp() (WebApp) {
	app := WebApp{}
	app.Init()

	ctrl := NewController()
	app.RegisterController(ctrl)

	app.BuildRouter()

	return app
}

func Start(listener net.Listener) {
	app := NewWebApp()
	app.Start(listener)
}



