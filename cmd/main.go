package main

import (
	"fmt"

	"github.com/minormending/coin-prices/clients"
)

func main() {
	server := clients.HTTPCrypto{}
	price, err := clients.CoinDeskPrice(&server, clients.CoinDeskBitcoinID, "USD")
	if err != nil {
		panic(err)
	}
	fmt.Println(price)

	price, err = clients.BlockChainPrice(&server, clients.BlockChainBitcoinID, "USD")
	if err != nil {
		panic(err)
	}
	fmt.Println(price)
}
