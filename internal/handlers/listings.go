package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/mbocsi/goapi-demo/api"
)

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
			h.handleGet(res)
		case "POST":
			h.handlePost(res, req)
		default:
			http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
		}
		return
	}
	// TODO:
	http.Error(res, "Not implemented", http.StatusNotImplemented)
}

func (h *ListingsHandler) handleGet(res http.ResponseWriter) {
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
}

func (h *ListingsHandler) handlePost(res http.ResponseWriter, req *http.Request) {
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
}
