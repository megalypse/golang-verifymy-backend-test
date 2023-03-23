package controllers

import (
	"net/http"
)

type ParsedRequest[T any] struct {
	Body   *T
	Params map[string]string
}

type HttpResponse struct {
	HttpStatus int    `json:"http_status"`
	Message    string `json:"message"`
	Content    any    `json:"content"`
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
