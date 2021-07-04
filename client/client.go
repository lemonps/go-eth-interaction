package client

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func (client *ethclient.Client) connect() (*ethclient.Client, error) {
	client, err := ethclient.Dial("https://rpc-testnet.bitkubchain.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("We are connected!")
	_ = client // we'll use this in the upcoming sections

	return client, nil
}
