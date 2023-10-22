package crypto_handler

// Get prices
type Prices struct {
	Prices map[string]float64
}

// Get history
type HistoricalData struct {
	TimeFrom int
	TimeTo   int
	Data     []OHLCVItem
}

type OHLCVItem struct {
	Time       int
	High       float64
	Low        float64
	Open       float64
	Close      float64
	VolumeFrom float64
	VolumeTo   float64
}

// Get difference
type PriceDifference struct {
	Diff         float64 `json:"diff"`
	DiffPercents float64 `json:"percents"`
}
