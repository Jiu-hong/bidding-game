const Bidding = artifacts.require('Bidding');
const { time } = require("@openzeppelin/test-helpers");

  
contract('Test Bidding', (accounts) => {
    it('should be able to bid', async () => {
        const contract_address = "0x0290FB167208Af455bB137780163b7B7a9a10C16"
   web3.eth.subscribe('logs', {
    address: contract_address,
    topics: [web3.utils.keccak256("HighestBidIncreased(address,uint256,uint256)")]
  }, function(error, result){
    if (!error)
        console.log(result);
  })
        });
});