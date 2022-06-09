const Bidding = artifacts.require('Bidding');
const { time } = require("@openzeppelin/test-helpers");

  
contract('Test Bidding', (accounts) => {
    it('should be able to bid', async () => {
        const BiddingInstance = await Bidding.new();
        assert.equal(await (await BiddingInstance.getHighestBid()).toString(),0,"initial highest bid should be 0");
        assert.equal(await (await BiddingInstance.getCommission()).toString(), 0, "initial commission should be 0");
        assert.equal(await (await BiddingInstance.gettotalRewards()).toString(), 0, "initial total rewards should be 0");

        //   bid
        await BiddingInstance.bid({ from: accounts[1], value: 2000 }) 

        assert.equal(await (await BiddingInstance.getHighestBid()).toString(),2000," highest bid should be 2000");
        assert.equal(await (await BiddingInstance.getCommission()).toString(), 100, " commission should be 100");
        assert.equal(await (await BiddingInstance.gettotalRewards()).toString(), 1900, " total rewards should be 1900")
    });
    
    it('should be able to end auction', async () => {
        const BiddingInstance = await Bidding.new();
        // first bid
        await BiddingInstance.bid({ from: accounts[1], value: web3.utils.toWei("10") })   // auction end time increases 10 minutes
        // second bid
        await BiddingInstance.bid({ from: accounts[2], value: web3.utils.toWei("20") })   // auction end time increases 10 minutes
        
        const total_rewards = web3.utils.fromWei(await BiddingInstance.gettotalRewards());
        assert.equal(web3.utils.fromWei(await BiddingInstance.getHighestBid()),20," highest bid should be 20 ethers");
        assert.equal(web3.utils.fromWei(await BiddingInstance.getCommission()), 1.5, " commission should be 1.5 ether");
        assert.equal(total_rewards, 28.5, " total rewards should be 28.5 ethers")

        await time.increase(4800) // 80 minutes == 4800 seconds

        // get balance before auction end
        const balance_account1_beforeEnd = web3.utils.fromWei(await web3.eth.getBalance(accounts[1]))
        const balance_account2_beforeEnd = web3.utils.fromWei(await web3.eth.getBalance(accounts[2]))
        // auction end
        await BiddingInstance.auctionEnd()

        // get balance after auction end
        const balance_account1_afterEnd = web3.utils.fromWei(await web3.eth.getBalance(accounts[1]))
        const balance_account2_afterEnd = web3.utils.fromWei(await web3.eth.getBalance(accounts[2]))
        assert.equal(balance_account1_beforeEnd, balance_account1_afterEnd, "balance of account1 shouldn't be changed")

        const account2_difference = balance_account2_afterEnd - balance_account2_beforeEnd
        assert.equal(account2_difference, total_rewards, "total rewards should be added to account2")

        // restarted game: HighestBid should be 0, totalRewards should be 0, however, commission should be accumulated 
        assert.equal(await (await BiddingInstance.getHighestBid()).toString(),0,"initial highest bid should be 0");
        assert.equal(web3.utils.fromWei(await BiddingInstance.getCommission()), 1.5, " commission should be 1.5 ether");
        assert.equal(await (await BiddingInstance.gettotalRewards()).toString(), 0, "initial total rewards should be 0");
    });
    
    it('AuctionNotYetEnded: should revert end auction before time expired', async () => {
        // AuctionNotYetEnded
        const BiddingInstance = await Bidding.new();
        // first bid
        await BiddingInstance.bid({ from: accounts[1], value: web3.utils.toWei("10") })   // auction end time increases 10 minutes
        // second bid
        await BiddingInstance.bid({ from: accounts[2], value: web3.utils.toWei("20") })   // auction end time increases 10 minutes
        
        await time.increase(3800) // 80 minutes == 4800 seconds

        try {
            await BiddingInstance.auctionEnd();
            assert(false); 
        } catch (error) {
            assert(error); 
        }

    });

    it('AuctionAlreadyEnded: should revert when current auction has already ended', async () => {
        // AuctionAlreadyEnded
        const BiddingInstance = await Bidding.new();

        await time.increase(4900) // 80 minutes == 4800 seconds

        try {
            await BiddingInstance.bid({ from: accounts[2], value: web3.utils.toWei("9") }) 
            assert(false); 
        } catch (error) {
            assert(error); 
        }

    });
    
    it('BidNotHighEnough: should revert when Bid Not High Enough', async () => {
        // BidNotHighEnough
        const BiddingInstance = await Bidding.new();

        await BiddingInstance.bid({ from: accounts[1], value: web3.utils.toWei("10") }) 

        try {
            await BiddingInstance.bid({ from: accounts[2], value: web3.utils.toWei("9") }) 
            assert(false); 
        } catch (error) {
            assert(error); 
        }

    });

});