package handlers

import (
	"net/http"

	"github.com/mbocsi/goapi-demo/api"
)

type UsersHandler struct {
	userService api.UserService
}

func NewUsersHandler(s api.UserService) *UsersHandler {
	return &UsersHandler{userService: s}
}

// TODO:
func (h *UsersHandler) serveHTTP(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "Not implemented", http.StatusNotImplemented)
}
