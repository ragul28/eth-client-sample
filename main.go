package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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

	createWallet()
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

func createWallet() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
}
