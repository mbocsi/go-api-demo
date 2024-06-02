package api

type Listing struct {
	Id     int     `json:"id"`
	UserId string  `json:"userId"`
	Name   string  `json:"name"`
	Price  float64 `json:"askingPrice"`
}

type ListingService interface {
	Listing(id int) (*Listing, error)
	Listings() ([]Listing, error)
	Create(listing *Listing) error
	Delete(id int) error
}

type ListingRepository interface {
	Find(id int) (*Listing, error)
	FindAll() ([]Listing, error)
	Create(listing *Listing) error
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
