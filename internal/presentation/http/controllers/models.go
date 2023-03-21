package controllers

import "net/http"

type ParsedRequest[T any] struct {
	Body   *T
	Params map[string]string
}

type HttpResponse struct {
	HttpStatus int
	Message    string
	Content    any
}

type BaseController interface {
	GetHandlers() []RouteDefinition
}

type RouteDefinition struct {
	Method       string
	Route        string
	HandlingFunc http.HandlerFunc
}
