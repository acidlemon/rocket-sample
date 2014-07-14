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

	app.AddRoute(
		"/*all",
		(&Controller{}).Wildcard,
		&rocket.View{},
	)
	app.AddRoute(
		"/",
		(&Controller{}).TopPage,
		&rocket.View{},
	)
	app.AddRoute(
		"/signin",
		(&Controller{}).Signin,
		&rocket.View{},
	)

	app.AddRoute(
		"/xslate",
		(&Controller{}).Xslate,
		rocket.NewXslateView(),
	)


	app.BuildRouter()

	return app
}

func Start(listener net.Listener) {
	app := NewWebApp()
	app.Start(listener)
}



