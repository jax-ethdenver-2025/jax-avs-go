// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

interface IRewardPool {
    function initialize(bytes32 _hash) external payable;
    function distributeRewards(string[] calldata nodeIds, uint256[] calldata scores) external;
}
