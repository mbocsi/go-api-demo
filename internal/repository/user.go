package repository

import (
	"slices"

	"github.com/mbocsi/goapi-demo/api"
)

type userRepository struct {
	DB []api.User
}

func NewUserRepository(db []api.User) api.UserRepository {
	return &userRepository{DB: db}
}

func (u *userRepository) Find(id string) (*api.User, error) {
	idx := slices.IndexFunc(u.DB, func(u api.User) bool { return u.Id == id })
	return &u.DB[idx], nil
}

func (u *userRepository) Create(user *api.User) error {
	u.DB = append(u.DB, *user)
	return nil
}

func (u *userRepository) Delete(id string) error {
	idx := slices.IndexFunc(u.DB, func(u api.User) bool { return u.Id == id })
	u.DB = append(u.DB[:idx], u.DB[idx+1:]...)
	return nil
}
