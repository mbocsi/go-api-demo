package repository

import "github.com/mbocsi/goapi-demo/api"

type listingRepository struct {
	DB []api.Listing
}

func NewListingRepository(db []api.Listing) api.ListingRepository {
	return &listingRepository{DB: db}
}

// TODO:
func (l *listingRepository) Find(id int) (*api.Listing, error) {
	return &api.Listing{}, nil
}

func (l *listingRepository) FindAll() ([]api.Listing, error) {
	return l.DB, nil
}

// TODO:
func (l *listingRepository) Create(listing *api.Listing) error {
	return nil
}

// TODO:
func (l *listingRepository) Delete(id int) error {
	return nil
}
