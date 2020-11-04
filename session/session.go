package session

import (
	"context"
	"net/http"
)

type sessionKeyTyp string

const (
	SessionKey sessionKeyTyp = "SESSION"
)

type ContentType string

const (
	ContentTypeJSON ContentType = "application/json"
	ContentTypeHTML ContentType = "text/html"
)

type Session struct {
	context.Context

	Request        *http.Request
	ResponseWriter http.ResponseWriter

	ContentType  ContentType
	ResponseData interface{}
}
