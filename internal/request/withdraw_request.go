package request

type WithdrawRequest struct {
	UserID string  `json:"user_id" binding:"required,uuid"`
	Amount float64 `json:"amount" binding:"required"`
}
