// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Script.sol";

import "../src/RewardPool.sol";

contract RewardPoolDeployerScript is Script {
    function run() public returns (address rewardPoolAddress) {
        vm.startBroadcast();
        return address(new RewardPool{salt: 0}());
    }
}
