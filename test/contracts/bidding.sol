// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.4;

uint constant extendedInterval = 10 minutes;
uint constant gameInterval = 60 minutes;

contract Bidding {
    uint public commission;
    uint public auctionEndTime;
    address payable public highestBidder;
    uint public highestBid;
    uint totalRewards; // total reward from all bidders

    // Events 
    event HighestBidIncreased(address indexed bidder, uint indexed amount, uint indexed commission);
    event AuctionEnded(address indexed winner, uint indexed amount);

    /// The auction has already ended.
    error AuctionAlreadyEnded();
    /// There is already a higher or equal bid.
    error BidNotHighEnough(uint highestBid);
    /// The auction has not ended yet.
    error AuctionNotYetEnded();

    constructor() {
        auctionEndTime = block.timestamp + gameInterval; //Start time
    }

    function getHighestBid() external view returns (uint) {
        return highestBid;
    }

    function getCommission() external view returns (uint){
        return commission;
    }

    function gettotalRewards() external view returns (uint){
        return totalRewards;
    }

    function bid() external payable {
        require(msg.value > 0);

        if (block.timestamp > auctionEndTime)
            {
                revert AuctionAlreadyEnded();
            }

        if (msg.value <= highestBid)
           {
                revert BidNotHighEnough(highestBid);
           }

        totalRewards += msg.value * 95 / 100 ;  // rewards to winner
        commission += msg.value * 5 / 100 ;     //commission for operation

        highestBidder = payable(msg.sender);
        highestBid = msg.value;
        auctionEndTime += extendedInterval; 
        emit HighestBidIncreased(msg.sender, msg.value, commission);

    }

    /// End the auction and send rewards
    /// to the winner.
    function auctionEnd() external {
        // 1. Conditions
        if (block.timestamp < auctionEndTime)
            {
                revert AuctionNotYetEnded();
            }

        // 2. Effects
        emit AuctionEnded(highestBidder, highestBid);

        // 3. Interaction
        highestBidder.transfer(totalRewards);
        _restart();
    }

    function _restart() internal {
        auctionEndTime = block.timestamp + gameInterval; // reset start time
        highestBidder = payable(address(0));             // reset highest bidder
        highestBid = 0;                                  // reset highest bid
        totalRewards = 0;                                // rest total rewards
    }
}