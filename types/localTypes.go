package localTypes

import userModel "lets-go/models/user"

type ProgrammingLanguage struct {
	Name string
}

type RegisterRequestData struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type RegisterResponseData struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    *userModel.User `json:"data"`
}

type LoginRequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseData struct {
	Status      string          `json:"status"`
	Message     string          `json:"message"`
	AccessToken string          `json:"accessToken"`
	Data        *userModel.User `json:"data"`
}

type ContextKey string

