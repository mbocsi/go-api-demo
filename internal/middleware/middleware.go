package middleware

import "net/http"

type ResponseHeader struct {
	handler     http.Handler
	headerName  string
	headerValue string
}

func NewResponseHeader(next http.Handler, headerName string, headerValue string) *ResponseHeader {
	return &ResponseHeader{next, headerName, headerValue}
}

func (rh *ResponseHeader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(rh.headerName, rh.headerValue)
	rh.handler.ServeHTTP(res, req)
}
