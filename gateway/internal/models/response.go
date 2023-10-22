package models

type Response struct {
	HttpCode  int    `json:"httpCode"`
	HaveError bool   `json:"haveError"`
	Error     *Error `json:"error"`
	Payload   *any   `json:"payload"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
