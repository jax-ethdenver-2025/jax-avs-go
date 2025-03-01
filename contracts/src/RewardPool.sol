// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {ERC20} from "solady/tokens/ERC20.sol";
import {Ownable} from "solady/auth/Ownable.sol";
import {Initializable} from "solady/utils/Initializable.sol";
import {Ed25519} from "./libraries/ED25519.sol";
import {IIncredibleSquaringTaskManager} from "./IIncredibleSquaringTaskManager.sol";

struct Signature {
    bytes32 k;
    bytes32 r;
    bytes32 s;
    bytes m;
}

contract RewardPool is Ownable, Initializable {
    IIncredibleSquaringTaskManager public avs;

    mapping(string => address) public nodeIdToBeneficiary;

    // Add storage for pool metadata
    bytes32 public contentHash;

    // mapping from peer id to wallet address that receives rewards
    mapping(string => bool) public peers;
    string[] public peerList;

    event PeerAdded(string indexed nodeId);
    event PeerRemoved(string indexed nodeId);
    // NOTE: we should probably not pass hash and do better
    //  indexing off chain
    event Deposit(uint256 amount, bytes32 hash);
    event RewardDistributed(address indexed user, uint256 reward);

    /* initializer / constructor */

    constructor() {
        _disableInitializers();
    }

    receive() external payable {}

    function initialize(bytes32 _hash) external payable initializer {
        contentHash = _hash;
        avs = IIncredibleSquaringTaskManager(msg.sender);

        emit Deposit(msg.value, _hash);
    }

    modifier onlyAVS() {
        require(msg.sender == address(avs), "Caller is not AVS");
        _;
    }

    /* state changing functions */

    // TODO: fix this / diagnose why we revert against rust bindings
    function enterPool(string memory nodeId, bytes32 k, bytes32 r, bytes32 s, address beneficiary) external {
        nodeIdToBeneficiary[nodeId] = beneficiary;
        bytes memory m = abi.encodePacked(beneficiary);

        require(!peers[nodeId], "Peer already in pool");
        require(bytes(nodeId).length > 0, "Invalid node ID");
        require(verify(k, r, s, m), "Invalid signature");
        peers[nodeId] = true;
        peerList.push(nodeId);
        emit PeerAdded(nodeId);
    }

    function deposit() external payable {
        require(msg.value > 0, "Invalid amount");
        emit Deposit(msg.value, contentHash);
    }

    function distributeRewards(string[] calldata nodeIds, uint256[] calldata scores) external onlyAVS {
        uint256 nodeLength = nodeIds.length;
        uint256 scoresLength = scores.length;

        require(nodeLength == scoresLength, "Invalid input array length");

        uint256 totalScore;
        for (uint256 i = 0; i < scoresLength; i++) {
            require(scores[i] < 10000, "Invalid score (must be less than 10000 bps)");
            totalScore += scores[i];
        }
        require(totalScore < 10000, "Invalid score inputs (total score must be less than 10000 bps)");

        uint256 contractBalance = address(this).balance;
        require(contractBalance > 0, "No funds to distribute");

        for (uint256 i = 0; i < nodeLength; i++) {
            address recipient = nodeIdToBeneficiary[nodeIds[i]];
            require(recipient != address(0), "Recipient address not set");

            uint256 amount = (contractBalance * scores[i]) / 10000;
            if (amount > 0) {
                (bool sent,) = payable(recipient).call{value: amount}("");
                require(sent, "Transfer failed");
            }
        }
    }

    /* view functions */

    function getPeers() external view returns (string[] memory) {
        return peerList;
    }

    function getHash() external view returns (bytes32) {
        return contentHash;
    }

    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }

    /* helper functions */

    function verify(bytes32 k, bytes32 r, bytes32 s, bytes memory m) public pure returns (bool) {
        return Ed25519.verify(k, r, s, m);
    }
}
