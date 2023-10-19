package coinmarket

import (
	"time"
)

type QuotesResponse struct {
	Status Status               `json:"status"`
	X      map[string]DataInner `json:"data"`
}

type Status struct {
	Timestamp time.Time
}

type DataInner struct {
	Id    int
	Name  string
	Quote Quote
}

type Quote struct {
	USD USD
}

type USD struct {
	Price float64
}
