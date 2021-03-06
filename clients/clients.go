package clients

// Coin is a unified identifier across stores for a coin.
type Coin string

const (
	// CoinBitcoin represents Bitcoin
	CoinBitcoin Coin = "BTC"
)

// CoinPrice represents the price of a coin at a given time.
type CoinPrice struct {
	Timestamp string
	Coin      Coin
	Currency  string
	Price     float64
}

// HTTPCoinServer provides an interface to mock
type HTTPCoinServer interface {
	Get(url string, j interface{}) (interface{}, error)
}
