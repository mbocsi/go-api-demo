package handlers

import (
	"net/http"
	"path"
	"strings"
)

// Helper function for seperating url head/tails
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

type App struct {
	ApiHandler *ApiHandler
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	if head == "api" {
		h.ApiHandler.ServeHTTP(res, req)
		return
	}
	http.Error(res, "Not found", http.StatusNotFound)
}

type ApiHandler struct {
	ListingsHandler *ListingsHandler
	UsersHandler    *UsersHandler
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

type ListingsHandler struct{}

func (h *ListingsHandler) serveHTTP(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "Not implemented", http.StatusNotImplemented)
}

type UsersHandler struct{}

func (h *UsersHandler) serveHTTP(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "Not implemented", http.StatusNotImplemented)
}
