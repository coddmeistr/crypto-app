package models

type Response struct {
	HttpCode  int
	HaveError bool
	Error     *Error `json:"error"`
	Payload   any    `json:"payload"`
}

type Error struct {
	Code    int
	Message string
}
