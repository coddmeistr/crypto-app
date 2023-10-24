package bitfinex

type ResponseInfo struct {
	Event    string `json:"event"`
	Channel  string `json:"channel"`
	ChanId   int64  `json:"chanId"`
	Key      string `json:"key"`
	Version  int64  `json:"version"`
	ServerId string `json:"serverId"`
}

type ResponseOHLCVUpdate struct {
	ChannelId int64
	Status    string
	Snapshot  *[]OHLCVItem
	Update    *OHLCVItem
}

type OHLCVItem struct {
	Timestamp int64
	Open      float64
	Close     float64
	High      float64
	Low       float64
	Volume    float64
}
