package main

import (
	"log"
	"net/http"

	"github.com/hitzhangjie/gwapp/app"
	"github.com/hitzhangjie/gwapp/ctrl"
	"github.com/hitzhangjie/gwapp/router"
	"github.com/hitzhangjie/gwapp/session"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

const (
	AppUIURI        = "/"
	serveAppUIPath  = "./static"
	serveAppUIAddr  = ":8000"
	serveAppSVRAddr = ":8080"

	ExtJSURI       = "/extjs"
	serveExtJSPath = "./extjs"
)

func main() {

	// serve app frontend
	go func() {

		http.Handle(AppUIURI, http.FileServer(http.Dir(serveAppUIPath)))
		http.Handle(ExtJSURI, http.FileServer(http.Dir(serveExtJSPath)))

		log.Printf("Listening on %s", serveAppUIAddr)
		err := http.ListenAndServe(serveAppUIAddr, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// serve app backend
	log.Printf("Listening on %s", serveAppSVRAddr)

	c := &HelloController{}
	app.Register("/Hello", router.MethodGet, c) // route /Hello to c.Hello
	app.Register("/Index", router.MethodGet, c) // route /Index to c.Index
	app.Run(app.WithAddress(serveAppSVRAddr))
}

type HelloController struct {
	ctrl.BaseController
}

func (ctrl *HelloController) Index(s *session.Session) error {
	ctrl.Template = "static/views/index.html"
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
