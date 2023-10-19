package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model
	Name  string  `json:"name"`
	ToUSD float64 `json:"tousd"`
}

type PriceDifference struct {
	Diff         float64 `json:"diff"`
	DiffPercents float64 `json:"percents"`
}
