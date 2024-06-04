package handlers

import (
	"net/http"
)

type ApiHandler struct {
	ListingsHandler http.Handler
	UsersHandler    http.Handler
}

func NewApiHandler(l *ListingsHandler, u *UsersHandler) *ApiHandler {
	return &ApiHandler{ListingsHandler: l, UsersHandler: u}
}

func (h *ApiHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	switch head {
	case "listings":
		h.ListingsHandler.ServeHTTP(res, req)
	case "users":
		h.UsersHandler.ServeHTTP(res, req)
	default:
		http.Error(res, "Not found", http.StatusNotFound)
	}
	return
}
