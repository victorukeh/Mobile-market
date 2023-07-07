package finance

import "github.com/victorukeh/mobile-market/pkg/v1/models"

type SetStatusDto struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CreateCashTypeDto struct {
	Success  bool             `json:"success"`
	Message  string           `json:"message"`
	CashType *models.CashType `json:"cashType"`
}
