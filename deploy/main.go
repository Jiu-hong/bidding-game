package main

import (
	"context"

	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"challenge/biddingpk"
	"challenge/events"
)

const GASLIMIT = 3000000
func main() {
	envs, err := godotenv.Read(".env")

	network := envs["NETWORK"]
	privatekey := envs["PRIVATEKEY"]

	privateKey, err := crypto.HexToECDSA(privatekey)
    if err != nil {
        log.Fatal(err)
    }


	client, err := ethclient.Dial(network)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	add:=crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(),add)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err!=nil {
		log.Fatal(err)
	}
	
	chainID, err := client.NetworkID(context.Background())
	if err!= nil {
		log.Fatal(err)
	}

	auth,err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err!= nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(GASLIMIT)
	auth.Nonce = big.NewInt(int64(nonce))
	a,tx, _, err := biddingpk.DeployBiddingpk(auth, client)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println("-----------------------------------")
	fmt.Println("contract address",a.Hex())
	fmt.Println("Transaction hash",tx.Hash().Hex())

	events.SubscribeEvents(a.Hex())
}
