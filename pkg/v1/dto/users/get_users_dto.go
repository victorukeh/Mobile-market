package users

import "github.com/victorukeh/mobile-market/pkg/v1/models"

type GetUser struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	User    models.User `json:"user"`
}

type GetUsers struct {
	Success bool           `json:"success"`
	Page    int            `json:"page"`
	Limit   int            `json:"limit"`
	Message string         `json:"message"`
	User    []*models.User `json:"user"`
}
