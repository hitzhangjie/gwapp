package main

import (
	"log"

	"github.com/hitzhangjie/gwapp/app"
	"github.com/hitzhangjie/gwapp/ctrl"
	"github.com/hitzhangjie/gwapp/router"
	"github.com/hitzhangjie/gwapp/session"
)

func init() {
	log.SetFlags(log.LstdFlags|log.Llongfile)
}

func main() {
	app.Register("/Hello", router.MethodGet, &HelloController{})
	app.Register("/Index", router.MethodGet, &HelloController{})
	app.Run()
}

type HelloController struct {
	ctrl.BaseController
}

func (ctrl *HelloController) Index(s *session.Session) error {
	ctrl.Template = "views/index.html"
	ctrl.Data = struct{ UserName string }{UserName: "zhijie"}

	dat, err := ctrl.Render()
	if err != nil {
		return err
	}

	s.ContentType = session.ContentTypeHTML
	s.ResponseData = dat
	return nil
}

func (ctrl *HelloController) Hello(s *session.Session) error {
	log.Printf("hello hello hello")

	s.ContentType = session.ContentTypeJSON
	s.ResponseData = map[string]interface{}{
		"retcode": 1,
		"retmsg":  "ok",
	}
	return nil
}
