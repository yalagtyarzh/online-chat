package middleware

import "net/http"

type ResponseStatusInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseStatusInterceptor(w http.ResponseWriter) *ResponseStatusInterceptor {
	return &ResponseStatusInterceptor{w, http.StatusOK}
}

func (rsi *ResponseStatusInterceptor) WriteHeader(code int) {
	rsi.statusCode = code
	rsi.ResponseWriter.WriteHeader(code)
}
