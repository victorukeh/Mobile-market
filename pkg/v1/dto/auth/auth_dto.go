package auth

import (
	"github.com/victorukeh/mobile-market/pkg/v1/models"
)

type LoginForm struct {
	Email    *string `json:"email" validate:"required,min=2,max=100"`
	Password *string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Success bool        `json:"success"`
	Token   string      `json:"token"`
	Message string      `json:"message"`
	User    models.User `json:"user"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type RegisterResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	User    models.User `json:"user"`
}
