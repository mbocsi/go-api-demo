package repository

import "github.com/mbocsi/goapi-demo/api"

type userRepository struct {
	DB []api.User
}

// TODO:
func NewUserRepository(db []api.User) api.UserRepository {
	return &userRepository{DB: db}
}

// TODO:
func (u *userRepository) Find(id string) (*api.User, error) {
	return &api.User{}, nil
}

// TODO:
func (u *userRepository) Create(user *api.User) error {
	return nil
}

// TODO:
func (u *userRepository) Delete(id string) error {
	return nil
}
