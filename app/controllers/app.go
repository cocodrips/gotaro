package controllers

import (
	"github.com/revel/revel"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}


func (c App) RqHandler(text string) revel.Result {
	fmt.Println(c.Params)
	fmt.Println(text)
	TestPost()
	return c.RenderJson(Test{hello()})
}
