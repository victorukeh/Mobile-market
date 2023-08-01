package finance

import "github.com/victorukeh/mobile-market/pkg/v1/models"

type CreateCashGroupDto struct {
	Success   bool                `json:"success"`
	Message   string              `json:"message"`
	CashGroup *[]models.CashGroup `json:"cashGroup"`
}
