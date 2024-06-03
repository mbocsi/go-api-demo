package handlers

import (
	"net/http"
)

type ApiHandler struct {
	ListingsHandler *ListingsHandler
	UsersHandler    *UsersHandler
}

func NewApiHandler(l *ListingsHandler, u *UsersHandler) *ApiHandler {
	return &ApiHandler{ListingsHandler: l, UsersHandler: u}
}

func (h *ApiHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	switch head {
	case "listings":
		h.ListingsHandler.serveHTTP(res, req)
	case "users":
		h.UsersHandler.serveHTTP(res, req)
	default:
		http.Error(res, "Not found", http.StatusNotFound)
	}
	return
}
