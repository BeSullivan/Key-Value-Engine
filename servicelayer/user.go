package servicelayer

import (
	"Kiwi/interfaces"
	"Kiwi/models"
)

type UserServiceLayerImpl struct {
	userDatalayer interfaces.UserDatalayer
}

func NewUserServiceLayer(userDL interfaces.UserDatalayer) interfaces.UserServicelayer {
	return &UserServiceLayerImpl{
		userDatalayer: userDL,
	}
}

func (userSL *UserServiceLayerImpl) RegisterService(username string, password string) (models.User, error) {

	return models.User{}, nil
}

func (userSL *UserServiceLayerImpl) LoginService(username string, password string) (models.User, error) {

	return models.User{}, nil
}
