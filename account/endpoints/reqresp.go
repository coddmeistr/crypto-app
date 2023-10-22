package endpoints

type BuyActivityRequest struct {
	Symbol string  `validate:"required"`
	Price  float64 `validate:"required,gt=0"`
}

type GetAccountBalanceResponse struct {
	AccountID uint    `json:"account_id"`
	USD       float64 `json:"usd"`
}

type CreateAccountRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required_without=Login,omitempty,email"`
}

type LoginResponse struct {
	UserID uint `json:"id"`
}

type GetAccountResponse struct {
	ID    uint   `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

type ActivityResponse struct {
	Symbol string  `json:"symbol"`
	Amount float64 `json:"amount"`
}
type GetActivitiesResponse struct {
	AccountID  uint               `json:"account_id"`
	Activities []ActivityResponse `json:"activities"`
}
