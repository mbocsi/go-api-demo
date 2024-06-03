package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

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
	head, tail := ShiftPath(req.URL.Path)
	if tail != "/" {
		http.Error(res, "Not found", http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(head)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid user id %q", head), http.StatusBadRequest)
		return
	}
	switch req.Method {
	case "GET":
		h.handleGetId(id, res)
	case "PUT":
		h.handlePut(id, res, req)
	case "DELETE":
		h.handleDelete(id, res)
	default:
		http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (h *ListingsHandler) handleGet(res http.ResponseWriter) {
	data, err := h.listingService.Listings()
	if err != nil {
		api.InternalErrorHandler(res)
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
		if err == api.InvalidArgumentError {
			http.Error(res, err.Error(), http.StatusUnprocessableEntity)
		} else {
			api.InternalErrorHandler(res)
		}
		return
	}
	res.WriteHeader(http.StatusCreated)
}

func (h *ListingsHandler) handleGetId(id int, res http.ResponseWriter) {
	data, err := h.listingService.Listing(id)
	if err != nil {
		if err == api.NotFoundError {
			http.Error(res, err.Error(), http.StatusNotFound)
		} else {
			api.InternalErrorHandler(res)
		}
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}

func (h *ListingsHandler) handlePut(id int, res http.ResponseWriter, req *http.Request) {
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	var l *api.Listing = new(api.Listing)
	err := d.Decode(l)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.listingService.Update(id, l)
	if err == nil {
		res.WriteHeader(http.StatusOK)
		return
	}
	if err != api.NotFoundError {
		if err == api.InvalidArgumentError {
			http.Error(res, err.Error(), http.StatusUnprocessableEntity)
		} else {
			api.InternalErrorHandler(res)
		}
		return
	}
	err = h.listingService.Create(l)
	if err != nil {
		api.InternalErrorHandler(res)
	}
	res.WriteHeader(http.StatusCreated)
}

func (h *ListingsHandler) handleDelete(id int, res http.ResponseWriter) {
	err := h.listingService.Delete(id)
	if err != nil {
		if err == api.NotFoundError {
			http.Error(res, err.Error(), http.StatusNotFound)
		} else {
			api.InternalErrorHandler(res)
		}
		return
	}
	res.WriteHeader(http.StatusOK)
}
