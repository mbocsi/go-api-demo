package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/mbocsi/goapi-demo/api"
)

type ListingsHandler struct {
	listingService api.ListingService
}

func NewListingsHandler(s api.ListingService) *ListingsHandler {
	return &ListingsHandler{listingService: s}
}

func (h *ListingsHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	head, tail := ShiftPath(req.URL.Path)
	if tail != "/" { // This is the last route handler
		http.Error(res, "Not found", http.StatusNotFound)
		return
	}
	// id not specified, return all listings
	if req.URL.Path == "/" {
		switch req.Method {
		case "GET":
			h.handleGet(res)
		case "POST":
			h.handlePost(res, req)
		case "OPTIONS":
			h.handleOptions(res)
		default:
			http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
		}
		return
	}
	// id was given in route, return single listing
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
	case "OPTIONS":
		h.handleOptionsId(res)
	default:
		http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (h *ListingsHandler) handleGet(res http.ResponseWriter) {
	data, err := h.listingService.Listings()
	if err != nil {
		api.InternalErrorHandler(res)
	}
	resData := api.ListingsResponse{Success: true, Listings: data}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(resData)
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

func (h *ListingsHandler) handleOptions(res http.ResponseWriter) {
	res.Header().Set("Allow", "GET, POST")
	if os.Getenv("GO_ENV") != "PROD" {
		res.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
	}
	res.WriteHeader(http.StatusNoContent)
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
	resData := api.ListingResponse{Success: true, Listing: data}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(resData)
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

func (ListingsHandler) handleOptionsId(res http.ResponseWriter) {
	res.Header().Set("Allow", "GET, PUT, DELETE")
	if os.Getenv("GO_ENV") != "PROD" {
		res.Header().Set("Access-Control-Allow-Methods", "GET, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
	}
	res.WriteHeader(http.StatusNoContent)
}
