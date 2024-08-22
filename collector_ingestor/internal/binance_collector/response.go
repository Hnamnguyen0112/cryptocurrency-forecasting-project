package binancecollector

type BinanceCommon struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
}
