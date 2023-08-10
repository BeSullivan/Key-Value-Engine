package controllerlayer

import (
	"Kiwi/interfaces"
	"Kiwi/models"
	"Kiwi/utils"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type UserControllerImpl struct {
	userService interfaces.UserServicelayer
}

func NewUserController(userServiceObject interfaces.UserServicelayer) *UserControllerImpl {
	return &UserControllerImpl{
		userService: userServiceObject,
	}
}

func (uc UserControllerImpl) RegisterHandler(c echo.Context) error {
	// Read Request Body
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: "Invalid JSON"})
	}

	jsonFormatValidationMsg, jsonFormatErr := utils.ValidateJsonFormat(jsonBody, "username", "password")
	if jsonFormatErr != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: jsonFormatValidationMsg})
	}

	userUsername := jsonBody["username"].(string)
	userPass := jsonBody["password"].(string)

	user, err := uc.userService.RegisterService(userUsername, userPass)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (uc UserControllerImpl) LoginHandler(c echo.Context) error {
	return nil
}
