package controllerlayer

import (
	"Kiwi/interfaces"
	"Kiwi/models"
	"Kiwi/utils"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	userService interfaces.UserServicelayer
}

func NewUserController(userServiceObject interfaces.UserServicelayer) *UserControllerImpl {
	return &UserControllerImpl{
		userService: userServiceObject,
	}
}

// @Summary		Register a new user
// @Description	Register a new user with the provided information
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			body	body		models.UserCreateRequest	true	"User registration details"
// @Success		201		{object}	models.UserResponse
// @Failure		400		{object}	models.ErrorResponseRegisterLogin
// @Failure		422		{object}	models.ErrorResponseRegisterLogin
// @Failure		500		{object}	models.ErrorResponseRegisterLogin
// @Router			/users/register [post]
func (uc *UserControllerImpl) RegisterHandler(c echo.Context) error {
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

// @Summary		User login
// @Description	Login with username and password
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			body	body		models.UserLoginRequest	true	"Login request body"
// @Success		200		{object}	models.UserResponse
// @Failure		400		{object}	models.ErrorResponseRegisterLogin
// @Failure		422		{object}	models.ErrorResponseRegisterLogin
// @Router			/users/login [post]
func (uc *UserControllerImpl) LoginHandler(c echo.Context) error {
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

	user, err := uc.userService.LoginService(userUsername, userPass)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{ResponseCode: 422, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, user)

}
