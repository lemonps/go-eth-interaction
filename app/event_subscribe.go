package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	token "go-interaction/token" // for demo

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// LogTransfer ..
// type LogTransfer interface {
// 	From common.Address
// 	To() common.Address
// 	Value() *big.Int
// }

// LogApproval ..
// type LogApproval struct {
// 	TokenOwner common.Address
// 	Spender    common.Address
// 	Tokens     *big.Int
// }

func main() {
	client, err := ethclient.Dial("wss://wss-testnet.bitkubchain.io")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0x0759a85e6a66Aba7BD2B7A91bbfed65d55094D76")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			// var transferEvent LogTransfer

			transferEvent, err := contractAbi.Unpack("Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Event: %s", transferEvent)

			// transferEvent = common.HexToAddress(vLog.Topics[1].Hex())
			// transferEvent = common.HexToAddress(vLog.Topics[2].Hex())

			// fmt.Printf("From: %s\n", transferEvent.From.Hex())
			// fmt.Printf("To: %s\n", transferEvent.To.Hex())
			// fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			// var approvalEvent LogApproval

			approvalEvent, err := contractAbi.Unpack("Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Event: %s", approvalEvent)
			// approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			// approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			// fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			// fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			// fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
		}

		fmt.Printf("\n\n")
	}
}
