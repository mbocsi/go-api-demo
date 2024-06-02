package database

import "github.com/mbocsi/goapi-demo/api"

type Database struct {
	Listings []api.Listing
	Users    []api.User
}

var listings []api.Listing = []api.Listing{
	{Id: 3626, Name: "Wireless mouse", Price: 15.5, UserId: "user_4363n62n67272"},
	{Id: 8373, Name: "Office chair", Price: 20.2, UserId: "user_426n276jn27j2"},
	{Id: 1205, Name: "PHYS 311 Textbook", Price: 35, UserId: "user_83gh3hgjmbmd"},
	{Id: 2986, Name: "M2 Macbook air", Price: 750, UserId: "user_d9249j094jg2"},
	{Id: 8623, Name: "Badger tickets", Price: 50.8, UserId: "user_m376m33in4"},
}

var users []api.User = []api.User{
	{Id: "user_d9249j094jg2", Username: "mbocsi", Email: "mb@example.com"},
	{Id: "user_426n276jn27j2", Username: "jdoe", Email: "jd@example.com"},
	{Id: "user_m376m33in4", Username: "jsmith", Email: "jsmith@example.com"},
}

var DB Database = Database{Listings: listings, Users: users}
