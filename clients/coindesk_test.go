package clients

import (
	"fmt"
	"strings"
	"testing"
)

type mockHTTP struct {
	responseObj interface{}
}

func (h mockHTTP) Get(url string, j interface{}) error {
	if h.responseObj == nil {
		return fmt.Errorf("unable to parse response from %s", url)
	}
	j = h.responseObj
	return nil
}

func TestCoinDeskPrice(t *testing.T) {
	cases := []struct {
		name     string
		coin     CoinDeskCoinType
		currency string
		resObj   interface{}
		price    CoinPrice
		errText  string
	}{
		{
			name:     "success",
			coin:     CoinDeskBitcoinID,
			currency: "USD",
			resObj: struct {
				Time string
			}{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			server := mockHTTP{
				responseObj: nil,
			}
			price, err := CoinDeskPrice(server, tc.coin, tc.currency)
			expectedSuccess := len(tc.errText) == 0
			if expectedSuccess {
				if err != nil {
					t.Errorf("expected success, but encountered: %s", err)
				} else if price == nil {
					t.Errorf("expected success, but got empty price: %v", price)
				} else if tc.price.Price != price.Price {
					t.Errorf("unexpected price, expected %f but got %f", tc.price.Price, price.Price)
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
