package clients

import (
	"fmt"
)

const coindeskURI = "https://api.coindesk.com/v1/%s/currentprice.json"

// CoinDeskCoinType denotes the 2 letter coin id
type CoinDeskCoinType string

const (
	// CoinDeskBitcoinID is the CoinDesk identifier for Bitcoin
	CoinDeskBitcoinID CoinDeskCoinType = "bpi"
)

// CoinDeskPrice returns the current price in the specified currency
func CoinDeskPrice(server HTTPCoinServer, coin CoinDeskCoinType, currency string) (*CoinPrice, error) {
	url := fmt.Sprintf(coindeskURI, coin)
	res, err := server.Get(url, coindeskResponse{})
	if err != nil {
		return nil, fmt.Errorf("unable to get price from CoinDesk: %v", err)
	}

	resWrapper := res.(*coindeskResponse)
	var coinID Coin
	var coinPrices map[string]coindeskPriceResponse
	switch coin {
	case CoinDeskBitcoinID:
		coinID = CoinBitcoin
		coinPrices = resWrapper.Bitcoin
	default:
		return nil, fmt.Errorf("unknown coin specified: %s", coin)
	}

	currencyPrice, ok := coinPrices[currency]
	if !ok {
		return nil, fmt.Errorf("unknown currency specified: %s", currency)
	}

	coinPrice := CoinPrice{
		Timestamp: resWrapper.Time["updatedISO"],
		Coin:      coinID,
		Currency:  currencyPrice.Currency,
		Price:     currencyPrice.Price,
	}
	return &coinPrice, nil
}

type coindeskResponse struct {
	Time       map[string]string                `json:"time"`
	Disclaimer string                           `json:"disclaimer"`
	CoinName   string                           `json:"chartName"`
	Bitcoin    map[string]coindeskPriceResponse `json:"bpi,omitempty"`
}

type coindeskPriceResponse struct {
	Currency            string  `json:"code"`
	CurrencySymbol      string  `json:"symbol"`
	CurrencyDescription string  `json:"description"`
	Price               float64 `json:"rate_float"`
}
