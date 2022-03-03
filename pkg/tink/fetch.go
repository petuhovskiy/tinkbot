package tink

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Result struct {
	Rate            float64
	FormattedString string
}

// https://api.tinkoff.ru/v1/currency_rates?from=USD&to=AED
func FetchExchangeRate(from, to string) (*Result, error) {
	requestURL := "https://api.tinkoff.ru/v1/currency_rates?from=" + from + "&to=" + to
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data ExchangeRateResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	for _, rate := range data.Payload.Rates {
		if rate.Category == "CUTransfersPro" {
			return convertRate(rate), nil
		}
	}

	return nil, errors.New("No rate found")
}

func convertRate(rate Rates) *Result {
	return &Result{
		Rate:            rate.Buy,
		FormattedString: fmt.Sprintf("1 %s = %v %s", rate.FromCurrency.Name, rate.Buy, rate.ToCurrency.Name),
	}
}
