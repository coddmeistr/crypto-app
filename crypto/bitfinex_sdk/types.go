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
	ChannelId int64        `json:"chanId"`
	Status    string       `json:"status"`
	Snapshot  *[]OHLCVItem `json:"snapshot"`
	Update    *OHLCVItem   `json:"update"`
}

type OHLCVItem struct {
	Timestamp int64   `json:"ts"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
}
