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
	serveStaticAddr  = ":8000"
	serveDynamicAddr = ":8080"

	serveExtJSPath = "./extjs"
)

func main() {

	// serve static site, on *:3000
	go func() {
		fs := http.FileServer(http.Dir(serveExtJSPath))
		http.Handle("/", fs)

		log.Printf("Listening on %s", serveStaticAddr)
		err := http.ListenAndServe(serveStaticAddr, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// serve dynamic request
	log.Printf("Listening on %s", serveDynamicAddr)
	app.Register("/Hello", router.MethodGet, &HelloController{})
	app.Register("/Index", router.MethodGet, &HelloController{})
	app.Run(app.WithAddress(serveDynamicAddr))
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
