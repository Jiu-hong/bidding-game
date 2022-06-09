package webapi

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"challenge/biddingpk"
)

const GASLIMIT = 3000000

func AuctionEnd(_privateKey string ,_contractAddress string) (string, error){
	envs, err := godotenv.Read(".env")
	network := envs["NETWORK"]

	client, err := ethclient.Dial(network)
	if err != nil {
		return "", err
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(_privateKey)
    if err != nil {
        return "", err
    }

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err!=nil {
		return "", err
	}
	
	chainID, err := client.NetworkID(context.Background())
	if err!= nil {
		return "", err
	}

	// contract address got from `go run deploy/main.go`
	contract_address := common.HexToAddress(_contractAddress)

    t, err := biddingpk.NewBiddingpk(contract_address, client)
	if err!= nil {
		return "", err
	}

	tx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err!= nil {
		return "", err
	}

	tx.GasPrice = gasPrice
	tx.GasLimit = uint64(GASLIMIT)

	transaction, err := t.AuctionEnd(tx)
	return transaction.Hash().Hex(), nil
}