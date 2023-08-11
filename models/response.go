package models

type Response struct {
	ResponseCode uint16 `json:"responsecode"`
	Message      string `json:"message"`
}

type ErrorResponseRegisterLogin struct {
	ResponseCode int    `json:"responsecode"`
	Message      string `json:"message"`
}

type UserResponse struct {
	ID       uint   `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
}

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
