package models

type PriceDifference struct {
	Diff         float64 `json:"diff"`
	DiffPercents float64 `json:"percents"`
}
