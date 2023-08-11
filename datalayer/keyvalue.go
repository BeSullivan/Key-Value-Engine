package datalayer

import (
	"Kiwi/interfaces"
	"Kiwi/models"

	"gorm.io/gorm"
)

type KiwiDataLayerImpl struct {
	db *gorm.DB
}

func NewKiwiDataLayer(db *gorm.DB) interfaces.KiwiDatalayer {
	return &KiwiDataLayerImpl{
		db: db,
	}
}

func (kiwiDL *KiwiDataLayerImpl) Set(key string, value string) (models.KeyValue, error) {
	return models.KeyValue{}, nil
}

func (kiwiDL *KiwiDataLayerImpl) Get(key string) models.KeyValue {
	return models.KeyValue{}
}

func (kiwiDL *KiwiDataLayerImpl) GetAll() ([]models.KeyValue, error) {
	return nil, nil
}

func (kiwiDL *KiwiDataLayerImpl) Delete(key string) error {
	return nil
}
