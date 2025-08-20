package param

import (
	"time"
)

type TransactionUpdateRequest struct {
	Type          *string    `json:"type"`
	Amount        *float64   `json:"amount"`
	CategoryID    *string    `json:"category_id"`
	Description   *string    `json:"description"`
	TransactionAt *time.Time `json:"transaction_at"`
}
