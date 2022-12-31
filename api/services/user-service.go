package services

import (
	"log"

	"github.com/bj-budhathoki/golang-api/api/repository"
	"github.com/bj-budhathoki/golang-api/dtos"
	"github.com/bj-budhathoki/golang-api/model"
	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dtos.UserUpdateDTOS) model.User
	Profile(userId string) model.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (s *userService) Update(user dtos.UserUpdateDTOS) model.User {
	userToUpdate := model.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	updatedUser := s.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (s *userService) Profile(userID string) model.User {
	return s.userRepository.ProfileUser(userID)
}
