package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const blockchainURI = "https://api.blockchain.com/v3/exchange/tickers/%s-%s"

// BlockChainCoinType denotes the 2 letter coin id
type BlockChainCoinType string

const (
	// BlockChainBitcoinID is the BlockChain identifier for Bitcoin
	BlockChainBitcoinID BlockChainCoinType = "BTC"
)

// BlockChainPrice returns the current price in the specified currency
func BlockChainPrice(coin BlockChainCoinType, currency string) (*CoinPrice, error) {
	currency = strings.ToUpper(currency)
	url := fmt.Sprintf(blockchainURI, coin, currency)
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to get price from Blockchain: %v", err)
	}
	defer res.Body.Close()

	type wrapperJSON struct {
		Symbol string  `json:"symbol"`
		Price  float64 `json:"last_trade_price"`
	}

	var resWrapper wrapperJSON
	if err := json.NewDecoder(res.Body).Decode(&resWrapper); err != nil {
		return nil, fmt.Errorf("unable to parse response: %s", err)
	}

	var coinID Coin
	switch coin {
	case BlockChainBitcoinID:
		coinID = CoinBitcoin
	default:
		return nil, fmt.Errorf("unknown coin specified: %s", coin)
	}

	coinPrice := CoinPrice{
		Timestamp: time.Now().String(),
		Coin:      coinID,
		Currency:  currency,
		Price:     resWrapper.Price,
	}
	return &coinPrice, nil
}
