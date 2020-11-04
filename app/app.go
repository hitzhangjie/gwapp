package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"git.code.oa.com/gw/gwapp/router"
	"git.code.oa.com/gw/gwapp/session"
)

const (
	defaultAddress = ":8080"
)

// Run 运行应用实例
func Run(opts ...Option) error {
	appOpts := &options{
		address: defaultAddress,
	}

	for _, o := range opts {
		o(appOpts)
	}

	paths := router.RouteMappings()

	for path, fn := range paths {
		// TODO 使这里的映射逻辑更健壮点
		vals := strings.Split(path, "_")

		http.HandleFunc(vals[0], func(w http.ResponseWriter, r *http.Request) {

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			err := r.ParseForm()
			if err != nil {
				log.Printf("http parse form error: %v", err)
				return
			}

			session := &session.Session{
				Context:        ctx,
				Request:        r,
				ResponseWriter: w,
			}

			err = fn(session)
			if err != nil {
				fmt.Fprintf(w, "error: %v", err)
				return
			}

			buf, err := json.Marshal(session.ResponsePayload)
			if err != nil {
				log.Printf("json marshal error: %v", err)
				return
			}

			n, err := session.ResponseWriter.Write(buf)
			if err != nil || n != len(buf) {
				log.Printf("http respond write %d bytes, error: %v", n, err)
			}
			log.Printf("request URI: %s params: %s, %v, response: %s\n", r.RequestURI, r.Form.Encode(), string(buf))
		})
	}

	return http.ListenAndServe(appOpts.address, nil)
}

type options struct {
	address string
}

// Option 应用选项
type Option func(*options)

// WithAddress 指定应用监听地址
func WithAddress(address string) Option {
	return func(opts *options) {
		opts.address = address
	}
}
