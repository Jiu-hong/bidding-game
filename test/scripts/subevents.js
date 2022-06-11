const Web3 = require('web3');

const main = async function () {
    const web3 = new Web3(Web3.givenProvider || "wss://kovan.infura.io/ws/v3/c1705cd194dc4d8b9bad0ad140770896");
    const contract_address = "0xC608BB03a4dFaCc47C9C5C1DDB91F9bB1053C168"
   web3.eth.subscribe('logs', {
    address: contract_address,
    topics: [web3.utils.keccak256("HighestBidIncreased(address,uint256,uint256)")]
  }, function(error, result){
    if (!error)
        console.log(result);
  });
    
  }
  
  main()