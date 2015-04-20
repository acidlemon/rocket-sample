package sample

import (
	"net"

	"gopkg.in/acidlemon/rocket.v2"
)

type WebApp struct {
	rocket.WebApp
}

func NewWebApp() WebApp {
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
