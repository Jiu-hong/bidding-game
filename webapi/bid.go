package webapi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"challenge/biddingpk"
)


func Bid(_privateKey string ,_contractAddress string, _biddingAmount string) (string, error) {
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

	biddingAmount := new(big.Int)
	amount, result := biddingAmount.SetString(_biddingAmount, 10)
	if !result {
		return "", err
	}
	tx.Value = amount

	transaction, err := t.Bid(tx)
	return transaction.Hash().Hex(), nil
}