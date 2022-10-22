package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "http://localhost:8545"
)

func main() {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	currentBlock(client)

	account := common.HexToAddress("0x21404C3Affb6e6E06FB0d2CADE323B18eA6F9018")
	getBalance(client, account)
}

func currentBlock(client *ethclient.Client) {
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(block.Number())
}

func getBalance(client *ethclient.Client, account common.Address) {
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
}
