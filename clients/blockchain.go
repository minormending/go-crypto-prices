package clients

import (
	"fmt"
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
func BlockChainPrice(server HTTPCoinServer, coin BlockChainCoinType, currency string) (*CoinPrice, error) {
	currency = strings.ToUpper(currency)
	url := fmt.Sprintf(blockchainURI, coin, currency)
	res, err := server.Get(url, blockchainResponse{})
	if err != nil {
		return nil, fmt.Errorf("unable to get price from Blockchain: %v", err)
	}

	resWrapper := res.(*blockchainResponse)
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

type blockchainResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"last_trade_price"`
}
