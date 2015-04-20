package sample

import (
	"fmt"
	"net/http"

	"gopkg.in/acidlemon/rocket.v2"
)

type Controller struct {
	*rocket.Controller
}

func NewController() *Controller {
	c := &Controller{rocket.NewController()}
	view := &rocket.View{}

	c.AddRoute("/*all", c.Wildcard, view)
	c.AddRoute("/", c.TopPage, view)
	c.AddRoute("/signin", c.Signin, view)

	return c
}

func (self *Controller) TopPage(c rocket.CtxData) {
	fmt.Println("toppage called")

	c.Res().StatusCode = http.StatusOK
	c.RenderText("you called / ")
}

func (self *Controller) Signin(c rocket.CtxData) {
	fmt.Println("signin called")
	c.Res().StatusCode = http.StatusOK

	value := rocket.RenderVars{
		"nishikawa": "ichirin",
		"ichinose":  "shogo",
		"acidlemon": "powawa",
	}
	c.Render("template.html", value)
}

func (self *Controller) Wildcard(c rocket.CtxData) {
	fmt.Println("wildcard called")
	c.Res().StatusCode = http.StatusOK

	value := rocket.RenderVars{}
	for k, v := range c.Params() {
		value[k] = v
	}
	value["_path"] = c.Args()

	c.RenderJSON(value)
}
