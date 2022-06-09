package webapi

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"challenge/biddingpk"
)

func GetHighestBid(_address string ,_contractAddress string) (string, error) {
	envs, err := godotenv.Read(".env")
	network := envs["NETWORK"]

	client, err := ethclient.Dial(network)
	if err != nil {
		return "",err
	}
	defer client.Close()

	// contract address got from `go run deploy/main.go`
	contract_address := common.HexToAddress(_contractAddress)

    t, err := biddingpk.NewBiddingpk(contract_address, client)
	if err!= nil {
		return "",err
	}

	address:=common.HexToAddress(_address)
	highestBid, err := t.GetHighestBid(&bind.CallOpts{
		From:address,
	})
	if err!= nil {
		return "",err
	}

	return highestBid.String(), nil
}
