package repository

import (
	"log"

	"github.com/bj-budhathoki/golang-api/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) model.User
	UpdateUser(user model.User) model.User
	VerfiyCredential(email string, password string) interface{}
	IsDuplicationEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) model.User
	ProfileUser(userId string) model.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db userConnection) CreateUser(user model.User) model.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db userConnection) VerfiyCredential(email string, password string) interface{} {
	var user model.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db userConnection) UpdateUser(user model.User) model.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}
func (db *userConnection) IsDuplicationEmail(email string) (tx *gorm.DB) {
	var user model.User
	return db.connection.Where("email = ?", email).Take(&user)
}
func (db *userConnection) FindByEmail(email string) model.User {
	var user model.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}
func (db *userConnection) ProfileUser(userID string) model.User {
	var user model.User
	db.connection.Find(&user, userID)
	return user
}
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to has a password")
	}
	return string(hash)
}
