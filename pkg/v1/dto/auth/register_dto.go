package auth

import (
	"github.com/victorukeh/mobile-market/pkg/v1/models"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type RegisterResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	User    models.User `json:"user"`
}
