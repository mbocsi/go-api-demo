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

func NewApp(h *ApiHandler) *App {
	return &App{ApiHandler: h}
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
