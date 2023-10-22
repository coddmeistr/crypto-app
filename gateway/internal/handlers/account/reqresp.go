package account_handler

type GetAccountResponse struct {
	ID    uint   `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

type GetAccountBalanceResponse struct {
	AccountID uint    `json:"account_id"`
	USD       float64 `json:"usd"`
}
