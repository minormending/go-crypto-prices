package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const coindeskURI = "https://api.coindesk.com/v1/%s/currentprice.json"

// CoinDeskCoinType denotes the 2 letter coin id
type CoinDeskCoinType string

const (
	// CoinDeskBitcoinID is the CoinDesk identifier for Bitcoin
	CoinDeskBitcoinID CoinDeskCoinType = "bpi"
)

// CoinDeskPrice returns the current price in the specified currency
func CoinDeskPrice(coin CoinDeskCoinType, currency string) (*CoinPrice, error) {
	url := fmt.Sprintf(coindeskURI, coin)
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to get price from CoinDesk: %v", err)
	}
	defer res.Body.Close()

	type priceJSON struct {
		Currency            string  `json:"code"`
		CurrencySymbol      string  `json:"symbol"`
		CurrencyDescription string  `json:"description"`
		Price               float64 `json:"rate_float"`
	}

	type wrapperJSON struct {
		Time       map[string]string    `json:"time"`
		Disclaimer string               `json:"disclaimer"`
		CoinName   string               `json:"chartName"`
		Bitcoin    map[string]priceJSON `json:"bpi,omitempty"`
	}

	var resWrapper wrapperJSON
	if err := json.NewDecoder(res.Body).Decode(&resWrapper); err != nil {
		return nil, fmt.Errorf("unable to parse response: %s", err)
	}

	var coinID Coin
	var coinPrices map[string]priceJSON
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
