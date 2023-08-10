package controller

import (
	"Kiwi/interfaces"
	"Kiwi/models"
	"Kiwi/utils"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type UserController struct {
	userService interfaces.UserServicelayer
}

func NewUserController(userServiceObject interfaces.UserServicelayer) *UserController {
	return &UserController{
		userService: userServiceObject,
	}
}

// @Summary Register a new user
// @Description Register a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param body body UserRegister true "User registration details"
// @Success 201 {object} UserResponse
// @Failure 400 {object} ErrorResponseRegisterLogin
// @Failure 422 {object} ErrorResponseRegisterLogin
// @Failure 500 {object} ErrorResponseRegisterLogin
// @Router /users/register [post]
func (uc UserController) RegisterHandler(c echo.Context) error {
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

func (uc UserController) LoginHandler(c echo.Context) error {
	return nil
}
