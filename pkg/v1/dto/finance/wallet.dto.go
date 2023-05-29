package finance

import "github.com/victorukeh/mobile-market/pkg/v1/models"

type CreateWalletResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Wallet  *models.Wallet `json:"wallet"`
}
