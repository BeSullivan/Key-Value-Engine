package datalayer

import (
	"Kiwi/interfaces"
	"Kiwi/models"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDataLayerImpl struct {
	db *gorm.DB
}

func NewUserDataLayer(db *gorm.DB) interfaces.UserDatalayer {
	return &UserDataLayerImpl{
		db: db,
	}
}

func (userDL *UserDataLayerImpl) Get(id int) (models.User, error) {

	var user models.User
	selectUser := userDL.db.Where("id = ?", id).Find(&user)
	if selectUser.Error != nil {
		msg := "failed to get user"
		return models.User{}, errors.New(msg)
	}

	return user, nil
}

func (userDL *UserDataLayerImpl) Create(username string, password string) (models.User, error) {
	err := userDL.CheckUnique(username)
	if err != nil {
		return models.User{}, errors.New("duplicate username")
	}

	var user models.User
	user.Username = username

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, errors.New("failed to hashing password")
	}

	user.Password = string(hash)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return models.User{}, errors.New("failed to create token")
	}
	user.Token = tokenString

	userDL.db.Save(&user)

	return user, nil
}

func (userDL *UserDataLayerImpl) Login(username string, password string) (models.User, error) {

	var user models.User
	userDL.db.Where("username = ?", username).Find(&user)

	// User not found
	if user.ID == 0 {
		return models.User{}, errors.New("invalid username")
	}

	// Invalid password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("wrong password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return models.User{}, errors.New("failed to create token")
	}
	user.Token = tokenString

	userDL.db.Save(&user)
	return user, nil
}

func (userDL *UserDataLayerImpl) CheckUnique(username string) error {

	var user models.User
	userDL.db.Where("username = ?", username).First(&user)

	if user.ID != 0 {
		return errors.New("input username has already been registered")
	}

	return nil
}
