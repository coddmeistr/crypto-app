package coinmarket

import "encoding/json"

type QuotesResponse struct {
	DataOuter DataOuter
}

type DataOuter struct {
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

type DynamicKeysAndData struct {
	Key   string          // This will hold the dynamic key
	Value json.RawMessage // This will hold the dynamic JSON data
	Data  DataInner       `json:"-"` // This will be populated after further unmarshalling
}
