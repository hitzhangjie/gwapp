package session

import (
	"context"
	"net/http"
)

type sessionKeyTyp string

const (
	SessionKey sessionKeyTyp = "SESSION"
)

type Session struct {
	context.Context

	Request         *http.Request
	ResponsePayload interface{}
	ResponseWriter  http.ResponseWriter
}
