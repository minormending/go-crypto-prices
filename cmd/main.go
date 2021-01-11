package main

import (
	"fmt"

	"github.com/minormending/coin-prices/sources"
)

func main() {
	price, err := sources.CoinDeskPrice(sources.CoinDeskBitcoinID, "USD")
	if err != nil {
		panic(err)
	}
	fmt.Println(price)

	price, err = sources.BlockChainPrice(sources.BlockChainBitcoinID, "USD")
	if err != nil {
		panic(err)
	}
	fmt.Println(price)
}
