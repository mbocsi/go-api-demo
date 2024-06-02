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

func (l *listingService) Create(listing *api.Listing) error {
	return l.listingRespository.Create(listing)
}

func (l *listingService) Delete(id int) error {
	return l.listingRespository.Delete(id)
}
