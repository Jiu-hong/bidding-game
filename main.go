package main

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	webapi "challenge/webapi"
)

func main() {
	router := gin.Default()
	router.POST("/bid/:privateKey/:contractAddress/:biddingAmount", bid)
	router.GET("/commission/:address/:contractAddress", getCommission)
	router.GET("/highestbid/:address/:contractAddress", getHighestBid)
	router.GET("/totalRewards/:address/:contractAddress", gettotalRewards)
	router.POST("/auctionend/:privateKey/:contractAddress", auctionEnd)

	router.Run("localhost:8080")
}

func bid(c *gin.Context) {
	privateKey:=  c.Param("privateKey")
	contractAddress := c.Param("contractAddress")
	biddingAmount := c.Param("biddingAmount")

	match_privateKey := regexp.MustCompile(`^[a-fA-F0-9]{64}$`).MatchString(privateKey)
	match_contractAddress := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`).MatchString(contractAddress)
	match_biddingAmount := regexp.MustCompile(`^[0-9]+$`).MatchString(biddingAmount)

	if !match_privateKey || !match_contractAddress || !match_biddingAmount {
		c.String(http.StatusInternalServerError, "invalid format of input. For example: <server>/bid/<privateKey>/<contractAddress>/<biddingAmount>")

	}

	txid,err := webapi.Bid(privateKey, contractAddress, biddingAmount)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		c.IndentedJSON(http.StatusOK, txid)
	}

}

func auctionEnd(c *gin.Context) {
	privateKey:=  c.Param("privateKey")
	contractAddress := c.Param("contractAddress")

	match_privateKey := regexp.MustCompile(`^[a-fA-F0-9]{64}$`).MatchString(privateKey)
	match_contractAddress := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`).MatchString(contractAddress)

	if !match_privateKey || !match_contractAddress {
		c.String(http.StatusInternalServerError, "invalid format of input. For example: <server>/auctionEnd/<privateKey>/<contractAddress>")
	}

	txid,err := webapi.AuctionEnd(privateKey, contractAddress)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		c.IndentedJSON(http.StatusOK, txid)
	}
}

func getCommission(c *gin.Context) {
	address:=  c.Param("address")
	contractAddress := c.Param("contractAddress")

	match_address := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`).MatchString(address)
	match_contractAddress := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`).MatchString(contractAddress)
	if !match_address || !match_contractAddress {
		c.String(http.StatusInternalServerError, "invalid format of input. For example: <server>/commission/<address>/<contractAddress>")
	}

	commission, err := webapi.GetCommission(address, contractAddress)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		c.IndentedJSON(http.StatusOK, commission)
	}	
}

func getHighestBid(c *gin.Context) {
	address:=  c.Param("address")
	contractAddress := c.Param("contractAddress")

	match_address := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`).MatchString(address)
	match_contractAddress := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`).MatchString(contractAddress)
	if !match_address || !match_contractAddress {
		c.String(http.StatusInternalServerError, "invalid format of input. For example: <server>/highestbid/<address>/<contractAddress>")
	}

	highestBid, err :=webapi.GetHighestBid(address, contractAddress)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		c.IndentedJSON(http.StatusOK, highestBid)
	}
	
}

func gettotalRewards(c *gin.Context) {
	address:=  c.Param("address")
	contractAddress := c.Param("contractAddress")

	match_address := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`).MatchString(address)
	match_contractAddress := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`).MatchString(contractAddress)
	if !match_address || !match_contractAddress {
		c.String(http.StatusInternalServerError, "invalid format of input. For example: <server>/totalRewards/<address>/<contractAddress>")
	}
	
	rewards, err := webapi.GettotalRewards(address, contractAddress)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		c.IndentedJSON(http.StatusOK, rewards)
	}
}


