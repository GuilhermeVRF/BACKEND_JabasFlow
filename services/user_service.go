package services

import (
	"jabas-flow/models"
	"jabas-flow/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) GetUsers() ([]models.User, error) {
	return userService.userRepository.GetUsers()
}

func (userService *UserService) GetUser(id int) (models.User, error) {
	return userService.userRepository.GetUser(id)
}

func (userService *UserService) CreateUser(user models.User) (models.User, error) {
	userId, err := userService.userRepository.InsertUser(user)
	if err != nil {
		return models.User{}, err
	}

	user.ID = int(userId)
	return user, nil
}

func (userService *UserService) UpdateUser(user models.User) (models.User, error) {
	return userService.userRepository.UpdateUser(user)
}