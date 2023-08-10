package datalayer

import (
	"Kiwi/interfaces"
	"Kiwi/models"

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

func (userDL *UserDataLayerImpl) Get(id int) ([]models.User, error) {
	return nil, nil
}

func (userDL *UserDataLayerImpl) Create(username string, password string) (string, models.User, error) {
	return "", models.User{}, nil
}

func (userDL *UserDataLayerImpl) Login(username, password string) (string, models.User, error) {
	return "", models.User{}, nil
}

func (userDL *UserDataLayerImpl) CheckUnique(username string) (string, error) {
	return "", nil
}
