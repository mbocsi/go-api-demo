package api

type Listing struct {
	Id     int
	UserId string
	Name   string
	Price  float64
}

type ListingsResponse struct {
	Code    int
	Results []Listing
}

type Error struct {
	Code    int
	Message string
}
