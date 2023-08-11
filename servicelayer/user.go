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

	user, err := userSL.userDatalayer.Create(username, password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (userSL *UserServiceLayerImpl) LoginService(username string, password string) (models.User, error) {

	user, err := userSL.userDatalayer.Login(username, password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
