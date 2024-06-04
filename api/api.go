package api

import (
	"errors"
	"net/http"
)

type Listing struct {
	Id     int     `json:"id"`
	UserId string  `json:"userId"`
	Name   string  `json:"name"`
	Price  float64 `json:"askingPrice"`
}

type ListingResponse struct {
	Success bool     `json:"success"`
	Listing *Listing `json:"listing"`
}

type ListingsResponse struct {
	Success  bool      `json:"success"`
	Listings []Listing `json:"listings"`
}

type ListingService interface {
	Listing(id int) (*Listing, error)
	Listings() ([]Listing, error)
	Create(listing *Listing) error
	Update(id int, listing *Listing) error
	Delete(id int) error
}

type ListingRepository interface {
	Find(id int) (*Listing, error)
	FindAll() ([]Listing, error)
	Create(listing *Listing) error
	Update(id int, listing *Listing) error
	Delete(id int) error
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserService interface {
	User(id string) (*User, error)
	Create(user *User) error
	Delete(id string) error
}

type UserRepository interface {
	Find(id string) (*User, error)
	Create(user *User) error
	Delete(id string) error
}

var NotFoundError = errors.New("Item not found")

var InvalidArgumentError = errors.New("Invalid arguments")

func InternalErrorHandler(w http.ResponseWriter) {
	http.Error(w, "An unexpected error occured.", http.StatusInternalServerError)
}
