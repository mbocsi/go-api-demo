package service

import "github.com/mbocsi/goapi-demo/api"

type listingService struct {
	listingRespository api.ListingRepository
}

func NewListingService(r api.ListingRepository) api.ListingService {
	return &listingService{listingRespository: r}
}

func (l *listingService) Listing(id int) (*api.Listing, error) {
	return l.listingRespository.Find(id)
}

func (l *listingService) Listings() ([]api.Listing, error) {
	return l.listingRespository.FindAll()
}

func (l *listingService) isValid(listing *api.Listing) bool {
	if listing.Name == "" {
		return false
	}
	if listing.Id <= 0 {
		return false
	}
	if listing.Price < 0 {
		return false
	}
	if listing.UserId == "" {
		return false
	}
	return true
}

func (l *listingService) Create(listing *api.Listing) error {
	if !l.isValid(listing) {
		return api.InvalidArgumentError
	}
	return l.listingRespository.Create(listing)
}

func (l *listingService) Update(id int, listing *api.Listing) error {
	listing.Id = id
	if !l.isValid(listing) {
		return api.InvalidArgumentError
	}
	return l.listingRespository.Update(id, listing)
}

func (l *listingService) Delete(id int) error {
	return l.listingRespository.Delete(id)
}
