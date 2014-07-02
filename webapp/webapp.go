package sample

import (
	"net"
	"net/http"
	"github.com/acidlemon/sixfold"
)

type WebApp struct {
	sixfold.WebApp
}

//func (app *WebApp) Start(listener net.Listener) {
//	http.Serve(listener, nil)
//}

func NewWebApp() (WebApp) {
	app := WebApp{}
	app.Init()


	// register routing
	obj := &Controller{}

	app.AddRoute("/", func(w http.ResponseWriter, r *http.Request){ obj.TopPage(w, r) } )
	app.AddRoute("/signin", func(w http.ResponseWriter, r *http.Request){ obj.Signin(w, r) } )
	app.AddRoute("/favicon.ico", func(w http.ResponseWriter, r *http.Request){ obj.Signin(w, r) } )

	app.BuildRouter()

	return app
}

func Start(listener net.Listener) {
	app := NewWebApp()
	app.Start(listener)
}





