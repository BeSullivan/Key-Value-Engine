package interfaces

import "Kiwi/models"

type (
	KiwiDatalayer interface {
		SetValueWithKey(key string, value string) (models.KeyValue, error)
		GetValueByKey(key string) models.KeyValue
		DeleteKey(key string) error
	}
)
