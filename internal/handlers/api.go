package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"path"
	"strings"

	"github.com/mbocsi/goapi-demo/api"
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

type ListingsHandler struct {
	listingService api.ListingService
}

func NewListingsHandler(s api.ListingService) *ListingsHandler {
	return &ListingsHandler{listingService: s}
}

func (h *ListingsHandler) serveHTTP(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		switch req.Method {
		case "GET":
			data, err := h.listingService.Listings()
			if err != nil {
				http.Error(
					res,
					fmt.Sprintf("An error occured when getting listings, %v", err),
					http.StatusInternalServerError,
				)
			}
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(data)
			return
		case "POST":
			d := json.NewDecoder(req.Body)
			d.DisallowUnknownFields()
			var l *api.Listing = new(api.Listing)
			err := d.Decode(l)
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
				return
			}
			l.Id = rand.Intn(10000)
			err = h.listingService.Create(l)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
			}
			res.WriteHeader(http.StatusCreated)
			return
		}
	}
	// TODO:
	http.Error(res, "Not implemented", http.StatusNotImplemented)
}

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
