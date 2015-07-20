package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}


type Test struct {
	Name string

}

func (c App) Index() revel.Result {
	return c.Render()
}


func (c App) Incoming() revel.Result {
	return c.RenderJson(Test{"Hello"})
}
