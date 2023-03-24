package httputils

import (
	"net/http"
)

type ParsedRequest[T any] struct {
	Body   *T
	Params map[string]string
}

type HttpResponse[T any] struct {
	HttpStatus int    `json:"http_status"`
	Message    string `json:"message"`
	Content    T      `json:"content"`
}

type BaseController interface {
	GetHandlers() []RouteDefinition
}

type RouteDefinition struct {
	Method        string
	Route         string
	HandlingFunc  http.HandlerFunc
	Unprotected   bool
	RequiredRoles []string
}

type Void struct{}
