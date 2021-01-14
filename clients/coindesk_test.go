package clients

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

type mockHTTP struct {
	responseObj *coindeskResponse
}

func (h mockHTTP) Get(url string, j interface{}) (interface{}, error) {
	if h.responseObj == nil {
		return nil, fmt.Errorf("unable to parse response from %s", url)
	}
	return h.responseObj, nil
}

func TestCoinDeskPrice(t *testing.T) {
	testTime := time.Now().String()
	templateTime := map[string]string{
		"updatedISO": testTime,
	}

	cases := []struct {
		name         string
		coindeskType CoinDeskCoinType
		coin         Coin
		currency     string
		price        float64
		errText      string
	}{
		{
			name:         "success",
			coindeskType: CoinDeskBitcoinID,
			coin:         CoinBitcoin,
			currency:     "USD",
			price:        32718,
		},
		{
			name:         "failure with bad coin",
			coindeskType: "NA",
			currency:     "USD",
			errText:      "unknown coin",
		},
		{
			name:         "failure with bad currency",
			coindeskType: CoinDeskBitcoinID,
			currency:     "ZZZ",
			errText:      "unknown currency",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			server := mockHTTP{
				responseObj: &coindeskResponse{
					Time: templateTime,
					Bitcoin: map[string]coindeskPriceResponse{
						"USD": coindeskPriceResponse{
							Currency: "USD",
							Price:    tc.price,
						},
						"EUR": coindeskPriceResponse{
							Currency: "EUR",
							Price:    tc.price,
						},
						"GBP": coindeskPriceResponse{
							Currency: "GBP",
							Price:    tc.price,
						},
					},
				},
			}
			price, err := CoinDeskPrice(server, tc.coindeskType, tc.currency)

			expectedSuccess := len(tc.errText) == 0
			if expectedSuccess {
				if err != nil {
					t.Errorf("expected success, but encountered: %s", err)
				} else if price == nil {
					t.Errorf("expected success, but got empty price: %v", price)
				} else if tc.price != price.Price {
					t.Errorf("unexpected price, expected %f but got %f", tc.price, price.Price)
				} else if testTime != price.Timestamp {
					t.Errorf("unexpected timestamp, expected %s but got %s", testTime, price.Timestamp)
				} else if tc.coin != price.Coin {
					t.Errorf("unexpected coin, expected %s but got %s", tc.coin, price.Coin)
				} else if tc.currency != price.Currency {
					t.Errorf("unexpected currency, expected %s but got %s", tc.currency, price.Currency)
				}
			} else {
				if err == nil {
					t.Error("expected failure, but did not get error")
				} else if !strings.Contains(err.Error(), tc.errText) {
					t.Errorf("unexpected error, expected %s but got %s", tc.errText, err)
				}
			}
		})
	}
}
