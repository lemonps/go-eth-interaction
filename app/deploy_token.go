package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	token "go-interaction/token"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// for demo
)

func main() {
	client, err := ethclient.Dial("https://rpc-testnet.bitkubchain.io")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA("9ec60dc570290852805f9fec193673ccb31ec4b244858b89a23f4e0438159a93")
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(30000000) // in units
	auth.GasPrice = gasPrice

	// Token valiable
	tokenName := "LEMON"
	tokenSymbol := "MON"
	totalSupply := big.NewInt(100000)
	address, tx, instance, err := token.DeployToken(auth, client, tokenName, tokenSymbol, totalSupply)
	if err != nil {
		panic(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}
