package controllers

type parsedRequest[T any] struct {
	Body   *T
	Params map[string]string
}

type httpResponse struct {
	HttpStatus int
	Message    string
	Content    any
}
