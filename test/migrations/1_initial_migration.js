const Bidding = artifacts.require("Bidding");

module.exports = async function (deployer) {

  await deployer.deploy(Bidding)
  const bidding = await Bidding.deployed()
  const contract_address = bidding.address

 web3.eth.subscribe('logs', {
  address: contract_address,
  topics: [web3.utils.keccak256("HighestBidIncreased(address,uint256,uint256)")]
}, function(error, result){
  if (!error)
      console.log(result);
});
  
}
  

