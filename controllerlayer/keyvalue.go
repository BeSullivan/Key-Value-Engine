package controllerlayer

import (
	"Kiwi/interfaces"

	"github.com/labstack/echo/v4"
)

type KiwiControllerImpl struct {
	kiwiService interfaces.KiwiServicelayer
}

func NewKiwiController(kiwiSL interfaces.KiwiServicelayer) *KiwiControllerImpl {
	return &KiwiControllerImpl{
		kiwiService: kiwiSL,
	}
}

func (kc *KiwiControllerImpl) AddValueHandler(c echo.Context) error {
	return nil
}

func (kc *KiwiControllerImpl) GetValueHandler(c echo.Context) error {
	return nil
}

func (kc *KiwiControllerImpl) ListHandler(c echo.Context) error {
	return nil
}
