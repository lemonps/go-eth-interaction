package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://rpc-testnet.bitkubchain.io")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0xC3c2716690232C15891D4D03590b1DC2D2c418F7")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034

	blockNumber := big.NewInt(1000)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	ethUnitBalance := new(big.Float)
	ethUnitBalance.SetString(pendingBalance.String())
	ethUnitValue := new(big.Float).Quo(ethUnitBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethUnitValue) // 25729324269165216042
}
