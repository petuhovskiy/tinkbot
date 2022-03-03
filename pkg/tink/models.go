package tink

type ExchangeRateResponse struct {
	TrackingID string  `json:"trackingId"`
	ResultCode string  `json:"resultCode"`
	Payload    Payload `json:"payload"`
}

type LastUpdate struct {
	Milliseconds int64 `json:"milliseconds"`
}

type FromCurrency struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	StrCode string `json:"strCode"`
}

type ToCurrency struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	StrCode string `json:"strCode"`
}

type Rates struct {
	Category     string       `json:"category"`
	FromCurrency FromCurrency `json:"fromCurrency"`
	ToCurrency   ToCurrency   `json:"toCurrency"`
	Buy          float64      `json:"buy"`
	Sell         float64      `json:"sell"`
}

type Payload struct {
	LastUpdate LastUpdate `json:"lastUpdate"`
	Rates      []Rates    `json:"rates"`
}
