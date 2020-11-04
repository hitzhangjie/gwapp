package main

import (
	"github.com/astaxie/beego"
)

func main() {
	controller := &MainController{}

	beego.Router("/hello", controller, "GET:Hello")
	beego.Run(":8080")
}

// MainController:
// The controller must implement ControllerInterface
// Usually we extends beego.Controller
type MainController struct {
	beego.Controller
}

// address: http://localhost:8080/hello GET
func (ctrl *MainController) Hello() {

	// beego-example/views/hello_world.html
	ctrl.TplName = "index.html"
	ctrl.Data["username"] = "zhijie"

	// don't forget this
	_ = ctrl.Render()
}
