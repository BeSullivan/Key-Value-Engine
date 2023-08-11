package servicelayer

import (
	"Kiwi/interfaces"
	"Kiwi/models"
)

type KiwiServicelayerImpl struct {
	kiwiDatalayer interfaces.KiwiDatalayer
}

func NewKiwiServiceLayer(kiwiDL interfaces.KiwiDatalayer) interfaces.KiwiServicelayer {
	return &KiwiServicelayerImpl{
		kiwiDatalayer: kiwiDL,
	}
}

func (kiwiSL *KiwiServicelayerImpl) SetValueByKey(key string, values string) error {
	return nil
}
func (kiwiSL *KiwiServicelayerImpl) GetValueByKey(key string) (string, error) {
	return "", nil

}
func (kiwiSL *KiwiServicelayerImpl) GetAllKeyValues() ([]models.KeyValue, error) {
	return nil, nil
}
