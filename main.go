package main

import (
	"log"

	"github.com/hitzhangjie/gwapp/app"
	"github.com/hitzhangjie/gwapp/ctrl"
	"github.com/hitzhangjie/gwapp/router"
	"github.com/hitzhangjie/gwapp/session"
)

func main() {
	err := router.Register("/Hello", router.MethodGet, &HelloController{})
	if err != nil {
		panic(err)
	}
	app.Run()
}

type HelloController struct {
	ctrl.BaseController
	//beego.Controller
}

// address: http://localhost:8080/hello GET
func (ctrl *HelloController) Hello(session *session.Session) error {
	log.Printf("hello hello hello")

	session.ResponsePayload = map[string]interface{}{
		"retcode": 1,
		"retmsg":  "ok",
	}

	// beego-example/views/hello_world.html
	//ctrl.TplName = "index.html"
	//ctrl.Data["username"] = "zhijie"
	//
	// don't forget this
	//_ = ctrl.Render()
	return nil
}
