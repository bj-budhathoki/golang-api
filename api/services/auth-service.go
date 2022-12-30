package services

import (
	"log"

	"github.com/bj-budhathoki/golang-api/api/repository"
	"github.com/bj-budhathoki/golang-api/dtos"
	"github.com/bj-budhathoki/golang-api/model"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerfiyCredential(email string, password string) interface{}
	CreateUser(user dtos.UserCreateDTOS) model.User
	FindByEmail(email string) model.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (s *authService) VerfiyCredential(email string, password string) interface{} {
	res := s.userRepository.VerfiyCredential(email, password)
	if v, ok := res.(model.User); ok {
		comparedPassword := comparedPassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}
func (s *authService) CreateUser(user dtos.UserCreateDTOS) model.User {
	userTocreate := model.User{}
	err := smapping.FillStruct(&userTocreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := s.userRepository.CreateUser(userTocreate)
	return res

}
func (s *authService) FindByEmail(email string) model.User {
	return s.userRepository.FindByEmail(email)
}
func (s *authService) IsDuplicateEmail(email string) bool {
	res := s.userRepository.IsDuplicationEmail(email)
	return !(res.Error == nil)
}
func comparedPassword(hashPassword string, password []byte) bool {
	byteHash := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
