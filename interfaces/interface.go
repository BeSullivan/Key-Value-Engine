package interfaces

import "Kiwi/models"

type (
	KiwiDatalayer interface {
		Set(key string, value string) (models.KeyValue, error)
		Get(key string) models.KeyValue
		GetAll() ([]models.KeyValue, error)
		Delete(key string) error
	}

	KiwiServicelayer interface {
		SetValueByKey(key string, values string) error
		GetValueByKey(key string) (string, error)
		GetAllKeyValues() ([]models.KeyValue, error)
	}

	UserDatalayer interface {
		Get(id int) (models.User, error)
		Create(username string, password string) (models.User, error)
		Login(username string, password string) (models.User, error)
		CheckUnique(username string) error
	}

	UserServicelayer interface {
		RegisterService(username string, password string) (models.User, error)
		LoginService(username string, password string) (models.User, error)
	}
)
