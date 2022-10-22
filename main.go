package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	currentBlock(client)
func currentBlock(client *ethclient.Client) {
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(block.Number())
}
}
