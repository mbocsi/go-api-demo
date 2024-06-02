package repository

import (
	"slices"

	"github.com/mbocsi/goapi-demo/api"
)

type listingRepository struct {
	DB []api.Listing
}

func NewListingRepository(db []api.Listing) api.ListingRepository {
	return &listingRepository{DB: db}
}

func (l *listingRepository) Find(id int) (*api.Listing, error) {
	idx := slices.IndexFunc(l.DB, func(l api.Listing) bool { return l.Id == id })
	return &l.DB[idx], nil
}

func (l *listingRepository) FindAll() ([]api.Listing, error) {
	return l.DB, nil
}

func (l *listingRepository) Create(listing *api.Listing) error {
	l.DB = append(l.DB, *listing)
	return nil
}

func (l *listingRepository) Delete(id int) error {
	idx := slices.IndexFunc(l.DB, func(l api.Listing) bool { return l.Id == id })
	l.DB = append(l.DB[:idx], l.DB[idx+1:]...)
	return nil
}
