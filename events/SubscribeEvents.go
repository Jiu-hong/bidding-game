package events

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	bidding "challenge/biddingpk"
)

type LogHighestBidIncreased struct {
	bidder common.Address
	amount *big.Int
    commission *big.Int
}

type LogAuctionEnded struct {
	winner common.Address
	amount *big.Int
}

var contractAbi, _= abi.JSON(strings.NewReader(string(bidding.BiddingpkABI)))

var logHighestBidIncreasedSig = []byte("HighestBidIncreased(address,uint256,uint256)")
var logAuctionEndedSig = []byte("AuctionEnded(address,uint256)")
var logHighestBidIncreasedSigHash = crypto.Keccak256Hash(logHighestBidIncreasedSig)

var logAuctionEndedSigHash = crypto.Keccak256Hash(logAuctionEndedSig)

func SubscribeEvents(_contract_address string) {
    envs, err := godotenv.Read(".env")
	networkws := envs["NETWORKWS"]

    fmt.Println("----------events logs-------------")

    client, err := ethclient.Dial(networkws)
    if err != nil {
        log.Fatal(err)
    }

    contractAddress := common.HexToAddress(_contract_address)
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
            readLog(vLog) 
        }
    }

}

func readLog(vLog types.Log) {
        fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
        fmt.Printf("Log Index: %d\n", vLog.Index)

        switch vLog.Topics[0].Hex() {
        case logHighestBidIncreasedSigHash.Hex():
            fmt.Printf("Log Name: HighestBidIncreased\n")

            var highestBidIncreasedEvent LogHighestBidIncreased

            err := contractAbi.UnpackIntoInterface (&highestBidIncreasedEvent, "HighestBidIncreased", vLog.Data)
            if err != nil {
                log.Fatal(err)
            }

            highestBidIncreasedEvent.bidder = common.HexToAddress(vLog.Topics[1].Hex())
        

            fmt.Printf("bidder: %s\n", highestBidIncreasedEvent.bidder.Hex())
            fmt.Printf("amount: %s\n", vLog.Topics[2].Hex())
            fmt.Printf("commission: %s\n", vLog.Topics[3].Hex())

        case logAuctionEndedSigHash.Hex():
            fmt.Printf("Log Name: auctionEndedEvent\n")

            var auctionEndedEvent LogAuctionEnded

            err := contractAbi.UnpackIntoInterface(&auctionEndedEvent, "AuctionEnded", vLog.Data)
            if err != nil {
                log.Fatal(err)
            }

            auctionEndedEvent.winner = common.HexToAddress(vLog.Topics[1].Hex())

            fmt.Printf("winner: %s\n", auctionEndedEvent.winner.Hex())
            fmt.Printf("amount: %s\n", vLog.Topics[2].Hex())
        }

        fmt.Printf("\n\n")
}