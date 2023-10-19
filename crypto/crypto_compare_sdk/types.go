package cryptocompare

type Prices struct {
	Prices map[string]float64
}

type HistoricalResponse struct {
	Response       string
	Message        string
	HistoricalData HistoricalData `json:"data"`
}

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
