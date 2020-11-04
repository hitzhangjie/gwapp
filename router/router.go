package router

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/hitzhangjie/gwapp/ctrl"
	"github.com/hitzhangjie/gwapp/session"
)

var (
	paths = map[string]HandleFunc{}
)

var (
	registerDuplicate         = errors.New("register duplicate path")
	registerInvalidController = errors.New("register invalid ctrl")
	registerPathNotFound      = errors.New("register path not found")
)

type HandleFunc func(*session.Session) error

type Method int

const (
	MethodGet     = 1 << iota // "GET"
	MethodHead                // "HEAD"
	MethodPost                // "POST"
	MethodPut                 // "PUT"
	MethodPatch               // "PATCH" // RFC 5789
	MethodDelete              // "DELETE"
	MethodConnect             // "CONNECT"
	MethodOptions             // "OPTIONS"
	MethodTrace               // "TRACE"
)

// Register TODO path -> camelcase -> 函数名?
func Register(path string, method Method, controller interface{}) error {
	if controller == nil {
		return fmt.Errorf("%w, nil controller", registerInvalidController)
	}

	orig := path
	path = makePath(path, method)
	if _, ok := paths[path]; ok {
		return registerDuplicate
	}

	rt := reflect.TypeOf(controller)
	rh := reflect.TypeOf((*ctrl.Controller)(nil)).Elem()

	if !rt.Implements(rh) {
		return fmt.Errorf("%w, not implements ctrl.Controller interface", registerInvalidController)
	}

	registered := false

	for i := 0; i < rt.NumMethod(); i++ {
		m := rt.Method(i)
		if m.Name != orig[1:] {
			continue
		}

		// MUST: 出参、入参数量、类型，必须匹配
		if m.Type.NumIn() != 2 || m.Type.In(1) != reflect.TypeOf((*session.Session)(nil)) {
			return fmt.Errorf("%w, method argument list invalid", registerInvalidController)
		}
		if m.Type.NumOut() != 1 {
			return fmt.Errorf("%w, method return value invalid", registerInvalidController)
		} else {
			rv := reflect.New(m.Type.Out(0))
			if _, ok := rv.Interface().(*error); !ok {
				return fmt.Errorf("%w, method return value invalid", registerInvalidController)
			}
		}

		fn := func(session *session.Session) error {
			in := []reflect.Value{
				reflect.ValueOf(controller),
				reflect.ValueOf(session),
			}
			vals := m.Func.Call(in)

			if vals[0].IsNil() {
				return nil
			}
			return vals[0].Interface().(error)
		}
		paths[path] = fn
		registered = true
		break
	}

	if !registered {
		return fmt.Errorf("%w, path not registered", registerInvalidController)
	}
	return nil
}

func Route(path string, method Method) (HandleFunc, error) {
	path = makePath(path, method)
	if v, ok := paths[path]; ok {
		return v, nil
	}
	return nil, registerPathNotFound
}

func makePath(path string, method Method) string {
	return fmt.Sprintf("%s_%d", path, method)
}

func RouteMappings() map[string]HandleFunc {
	return paths
}
