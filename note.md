##  cd work directory
```
cd Judy
```
##   Compile solidity 
```
solc --bin --abi contract/bidding.sol -o build
abigen --bin=build/Bidding.bin --abi=build/Bidding.abi --pkg=biddingpk --out=biddingpk/biddingpk.go
```

## Setting .env
```
cp .env_example .env
```
Then replace with your **NETWORK**, **NETWORKWS**(for events subscribing ) and **PRIVATEKEY**

(see example at bottom)

##   Deploy contract and subscribe events
```
go run deploy/main.go
```

*write down contract address. For example, 
`0x28464074e9b387636D11fa8d312101173e230F4e`*

##   Start server
```
go run main.go
```

##   Interact with server

get commission: GET

http://localhost:8080/commission/0x0aC86d7d69c4145F6892A1035e24A3B30CC42e9f/0x28464074e9b387636D11fa8d312101173e230F4e

get highestbid: GET

http://localhost:8080/highestbid/0x0aC86d7d69c4145F6892A1035e24A3B30CC42e9f/0x28464074e9b387636D11fa8d312101173e230F4e

get totalRewards: GET

http://localhost:8080/totalRewards/0x0aC86d7d69c4145F6892A1035e24A3B30CC42e9f/0x28464074e9b387636D11fa8d312101173e230F4e

bid: POST

http://localhost:8080/bid/6b71fd551da47acd89809f1ba299293bc54145e57bae3958f8cf4f8215563968/0x28464074e9b387636D11fa8d312101173e230F4e/100

auctionend: POST

http://localhost:8080/auctionend/6b71fd551da47acd89809f1ba299293bc54145e57bae3958f8cf4f8215563968/0x28464074e9b387636D11fa8d312101173e230F4e



for example:

*private key:
6b71fd551da47acd89809f1ba299293bc54145e57bae3958f8cf4f8215563968*

*address:
0x0aC86d7d69c4145F6892A1035e24A3B30CC42e9f*

*contract address
0x28464074e9b387636D11fa8d312101173e230F4e*