package service

import "github.com/mbocsi/goapi-demo/api"

type userService struct {
	userRepository api.UserRepository
}

func NewUserService(r api.UserRepository) api.UserService {
	return &userService{userRepository: r}
}

func (u *userService) User(id string) (*api.User, error) {
	return u.userRepository.Find(id)
}

func (u *userService) Create(user *api.User) error {
	return u.userRepository.Create(user)
}

func (u *userService) Delete(id string) error {
	return u.userRepository.Delete(id)
}
