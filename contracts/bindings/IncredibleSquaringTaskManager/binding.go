// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIncredibleSquaringTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IIncredibleSquaringTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSquaringTaskManagerTask struct {
	FileHash                  [32]byte
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IIncredibleSquaringTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSquaringTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	Providers          []common.Address
	Scores             []*big.Int
}

// IIncredibleSquaringTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSquaringTaskManagerTaskResponseMetadata struct {
	TaskRespondedBlock uint32
	HashOfNonSigners   [32]byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractIncredibleSquaringTaskManagerMetaData contains all meta data concerning the ContractIncredibleSquaringTaskManager contract.
var ContractIncredibleSquaringTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createPool\",\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllPools\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponse\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"providers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"scores\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPool\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pools\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueFileHash\",\"inputs\":[{\"name\":\"fileHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"components\":[{\"name\":\"fileHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"providers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"scores\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"components\":[{\"name\":\"fileHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"providers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"scores\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAggregator\",\"inputs\":[{\"name\":\"newAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGenerator\",\"inputs\":[{\"name\":\"newGenerator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccessfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usedHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AggregatorUpdated\",\"inputs\":[{\"name\":\"oldAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GeneratorUpdated\",\"inputs\":[{\"name\":\"oldGenerator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGenerator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"components\":[{\"name\":\"fileHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolCreated\",\"inputs\":[{\"name\":\"poolAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"balance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"providers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"scores\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b50604051620068f3380380620068f38339810160408190526200003591620001ea565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b5919062000231565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000133919062000231565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b3919062000231565b6001600160a01b031660e0525063ffffffff16610100525062000258565b6001600160a01b0381168114620001e757600080fd5b50565b60008060408385031215620001fe57600080fd5b82516200020b81620001d1565b602084015190925063ffffffff811681146200022657600080fd5b809150509250929050565b6000602082840312156200024457600080fd5b81516200025181620001d1565b9392505050565b60805160a05160c05160e05161010051616609620002ea600039600081816102d8015281816108b1015261381a01526000818161082d01526125b40152600081816105940152818161279601526131bb0152600081816105c8015281816129520152612afb0152600081816105fc01528181611383015281816122960152818161242e015261265101526166096000f3fe6080604052600436106102235760003560e01c806310d67a2f14610228578063136439dd1461024a5780631459457a1461026a578063171f1d5b1461028a5780631ad43189146102c65780631ce8a3e31461030f578063245a7bfc1461034f5780632cb223d51461037c5780632d89f6fc146103b757806331b36bd9146103e45780633563b0d114610411578063416c7e5e1461043e5780634a7c7e4b1461045e5780634d2b57fe1461047e5780634f739f74146104ab578063595c6a67146104d85780635ac86ab7146104ed5780635b16ebb71461051d5780635c1556621461054d5780635c975abb1461056d5780635df459461461058257806368304835146105b65780636d14a987146105ea5780636efb46361461061e578063715018a61461064c57806372d18e8d146106615780637afa1eed1461067c5780638866d1871461069c578063886f1195146106bc5780638889df3c146106dc5780638b00ce7c146106fc5780638da5cb5b14610719578063ac4afa381461072e578063aef18bf71461074e578063b98d09081461077e578063bb972b0414610798578063cefa7799146107b8578063cefdc1d4146107d8578063d88ff1f414610806578063df5cf7231461081b578063e25a54a91461084f578063ee2373e91461086f578063f2fde38b14610882578063f5c9899d146108a2578063f63c5bab146108d5578063f9120af6146108ea578063fabc1cbc1461090a578063ffb3959f1461092a575b600080fd5b34801561023457600080fd5b50610248610243366004614de7565b610957565b005b34801561025657600080fd5b50610248610265366004614e04565b610a13565b34801561027657600080fd5b50610248610285366004614e1d565b610b40565b34801561029657600080fd5b506102aa6102a5366004614ff3565b610c8f565b6040805192151583529015156020830152015b60405180910390f35b3480156102d257600080fd5b506102fa7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016102bd565b34801561031b57600080fd5b5061033f61032a366004615061565b60cc6020526000908152604090205460ff1681565b60405190151581526020016102bd565b34801561035b57600080fd5b5060cd5461036f906001600160a01b031681565b6040516102bd919061507e565b34801561038857600080fd5b506103a9610397366004615061565b60cb6020526000908152604090205481565b6040519081526020016102bd565b3480156103c357600080fd5b506103a96103d2366004615061565b60ca6020526000908152604090205481565b3480156103f057600080fd5b506104046103ff3660046150b5565b610df5565b6040516102bd91906151a3565b34801561041d57600080fd5b5061043161042c3660046151b6565b610f02565b6040516102bd9190615311565b34801561044a57600080fd5b50610248610459366004615332565b611381565b34801561046a57600080fd5b50610248610479366004614de7565b6114b7565b34801561048a57600080fd5b5061049e6104993660046153b5565b6114c8565b6040516102bd919061543d565b3480156104b757600080fd5b506104cb6104c6366004615498565b6115dd565b6040516102bd9190615591565b3480156104e457600080fd5b50610248611c98565b3480156104f957600080fd5b5061033f61050836600461565b565b606654600160ff9092169190911b9081161490565b34801561052957600080fd5b5061033f610538366004614de7565b60d16020526000908152604090205460ff1681565b34801561055957600080fd5b50610404610568366004615678565b611d52565b34801561057957600080fd5b506066546103a9565b34801561058e57600080fd5b5061036f7f000000000000000000000000000000000000000000000000000000000000000081565b3480156105c257600080fd5b5061036f7f000000000000000000000000000000000000000000000000000000000000000081565b3480156105f657600080fd5b5061036f7f000000000000000000000000000000000000000000000000000000000000000081565b34801561062a57600080fd5b5061063e61063936600461594c565b611f03565b6040516102bd929190615a0c565b34801561065857600080fd5b50610248612dac565b34801561066d57600080fd5b5060c95463ffffffff166102fa565b34801561068857600080fd5b5060ce5461036f906001600160a01b031681565b3480156106a857600080fd5b506102486106b7366004615a79565b612dc0565b3480156106c857600080fd5b5060655461036f906001600160a01b031681565b3480156106e857600080fd5b506102486106f7366004614e04565b60d355565b34801561070857600080fd5b5060c9546102fa9063ffffffff1681565b34801561072557600080fd5b5061036f6132f3565b34801561073a57600080fd5b5061036f610749366004614e04565b613302565b34801561075a57600080fd5b5061033f610769366004614e04565b60d26020526000908152604090205460ff1681565b34801561078a57600080fd5b5060975461033f9060ff1681565b3480156107a457600080fd5b506102486107b3366004615b1d565b61332c565b3480156107c457600080fd5b5060cf5461036f906001600160a01b031681565b3480156107e457600080fd5b506107f86107f3366004615b71565b613447565b6040516102bd929190615ba8565b34801561081257600080fd5b5061049e6135d9565b34801561082757600080fd5b5061036f7f000000000000000000000000000000000000000000000000000000000000000081565b34801561085b57600080fd5b5061024861086a366004615bc9565b61363b565b61036f61087d366004614e04565b613aad565b34801561088e57600080fd5b5061024861089d366004614de7565b613bc9565b3480156108ae57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006102fa565b3480156108e157600080fd5b506102fa606481565b3480156108f657600080fd5b50610248610905366004614de7565b613c3f565b34801561091657600080fd5b50610248610925366004614e04565b613c50565b34801561093657600080fd5b5061094a610945366004615061565b613da7565b6040516102bd9190615c50565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ce9190615ca2565b6001600160a01b0316336001600160a01b031614610a075760405162461bcd60e51b81526004016109fe90615cbf565b60405180910390fd5b610a1081613e39565b50565b60655460405163237dfb4760e11b81526001600160a01b03909116906346fbf68e90610a4390339060040161507e565b602060405180830381865afa158015610a60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a849190615cf7565b610aa05760405162461bcd60e51b81526004016109fe90615d14565b60665481811614610b145760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d707420604482015277746f20756e70617573652066756e6374696f6e616c69747960401b60648201526084016109fe565b60668190556040518181523390600080516020616514833981519152906020015b60405180910390a250565b600054610100900460ff1615808015610b605750600054600160ff909116105b80610b7a5750303b158015610b7a575060005460ff166001145b610bdd5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016109fe565b6000805460ff191660011790558015610c00576000805461ff0019166101001790555b610c0b866000613f30565b610c1485614008565b610c1d8461405a565b610c26836140ac565b60cf80546001600160a01b0319166001600160a01b0384161790558015610c87576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610cd757610cd7615d4a565b60200201518951600160200201518a60200151600060028110610cfc57610cfc615d4a565b60200201518b60200151600160028110610d1857610d18615d4a565b602090810291909101518c518d830151604051610d759a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c610d989190615d60565b9050610de7610db1610daa88846140fe565b8690614189565b610db9614211565b610ddd610dce85610dc86142d1565b906140fe565b610dd78c6142f2565b90614189565b886201d4c0614376565b909890975095505050505050565b606081516001600160401b03811115610e1057610e10614e8e565b604051908082528060200260200182016040528015610e39578160200160208202803683370190505b50905060005b8251811015610efb57836001600160a01b03166313542a4e848381518110610e6957610e69615d4a565b60200260200101516040518263ffffffff1660e01b8152600401610e8d919061507e565b602060405180830381865afa158015610eaa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ece9190615d82565b828281518110610ee057610ee0615d4a565b6020908102919091010152610ef481615db1565b9050610e3f565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f689190615ca2565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610faa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fce9190615ca2565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611010573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110349190615ca2565b9050600086516001600160401b0381111561105157611051614e8e565b60405190808252806020026020018201604052801561108457816020015b606081526020019060019003908161106f5790505b50905060005b87518110156113755760008882815181106110a7576110a7615d4a565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015611108573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111309190810190615dcc565b905080516001600160401b0381111561114b5761114b614e8e565b60405190808252806020026020018201604052801561119657816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816111695790505b508484815181106111a9576111a9615d4a565b602002602001018190525060005b815181101561135f576040518060600160405280876001600160a01b03166347b314e88585815181106111ec576111ec615d4a565b60200260200101516040518263ffffffff1660e01b815260040161121291815260200190565b602060405180830381865afa15801561122f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112539190615ca2565b6001600160a01b0316815260200183838151811061127357611273615d4a565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106112a1576112a1615d4a565b6020026020010151878f6040518463ffffffff1660e01b81526004016112c993929190615e5c565b602060405180830381865afa1580156112e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130a9190615e7b565b6001600160601b031681525085858151811061132857611328615d4a565b6020026020010151828151811061134157611341615d4a565b6020026020010181905250808061135790615db1565b9150506111b7565b505050808061136d90615db1565b91505061108a565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114039190615ca2565b6001600160a01b0316336001600160a01b0316146114ae5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527b391037b3103a3432903932b3b4b9ba393ca1b7b7b93234b730ba37b960211b608482015260a4016109fe565b610a108161459a565b6114bf6145e1565b610a10816140ac565b606081516001600160401b038111156114e3576114e3614e8e565b60405190808252806020026020018201604052801561150c578160200160208202803683370190505b50905060005b8251811015610efb57836001600160a01b031663296bb06484838151811061153c5761153c615d4a565b60200260200101516040518263ffffffff1660e01b815260040161156291815260200190565b602060405180830381865afa15801561157f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115a39190615ca2565b8282815181106115b5576115b5615d4a565b6001600160a01b03909216602092830291909101909101526115d681615db1565b9050611512565b6115e5614c9c565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015611625573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116499190615ca2565b9050611653614c9c565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90611683908b9089908990600401615eda565b600060405180830381865afa1580156116a0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116c89190810190615efa565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906116fa908b908b908b90600401615fb1565b600060405180830381865afa158015611717573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261173f9190810190615efa565b6040820152856001600160401b0381111561175c5761175c614e8e565b60405190808252806020026020018201604052801561178f57816020015b606081526020019060019003908161177a5790505b50606082015260005b60ff8116871115611ba9576000856001600160401b038111156117bd576117bd614e8e565b6040519080825280602002602001820160405280156117e6578160200160208202803683370190505b5083606001518360ff168151811061180057611800615d4a565b602002602001018190525060005b86811015611aa95760008c6001600160a01b03166304ec63518a8a8581811061183957611839615d4a565b905060200201358e8860000151868151811061185757611857615d4a565b60200260200101516040518463ffffffff1660e01b815260040161187d93929190615fd1565b602060405180830381865afa15801561189a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118be9190615fed565b90506001600160c01b0381166119615760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527b3132903932b3b4b9ba32b932b21030ba10313637b1b5b73ab6b132b960211b608482015260a4016109fe565b8a8a8560ff1681811061197657611976615d4a565b6001600160c01b03841692013560f81c9190911c600190811614159050611a9657856001600160a01b031663dd9846b98a8a858181106119b8576119b8615d4a565b905060200201358d8d8860ff168181106119d4576119d4615d4a565b9050013560f81c60f81b60f81c8f6040518463ffffffff1660e01b8152600401611a0093929190615e5c565b602060405180830381865afa158015611a1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a419190616016565b85606001518560ff1681518110611a5a57611a5a615d4a565b60200260200101518481518110611a7357611a73615d4a565b63ffffffff9092166020928302919091019091015282611a9281615db1565b9350505b5080611aa181615db1565b91505061180e565b506000816001600160401b03811115611ac457611ac4614e8e565b604051908082528060200260200182016040528015611aed578160200160208202803683370190505b50905060005b82811015611b6e5784606001518460ff1681518110611b1457611b14615d4a565b60200260200101518181518110611b2d57611b2d615d4a565b6020026020010151828281518110611b4757611b47615d4a565b63ffffffff9092166020928302919091019091015280611b6681615db1565b915050611af3565b508084606001518460ff1681518110611b8957611b89615d4a565b602002602001018190525050508080611ba190616033565b915050611798565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611bea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c0e9190615ca2565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611c41908b908b908e90600401616053565b600060405180830381865afa158015611c5e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c869190810190615efa565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81526001600160a01b03909116906346fbf68e90611cc890339060040161507e565b602060405180830381865afa158015611ce5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d099190615cf7565b611d255760405162461bcd60e51b81526004016109fe90615d14565b600019606681905560405190815233906000805160206165148339815191529060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611d8492919061607d565b600060405180830381865afa158015611da1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611dc99190810190615efa565b9050600084516001600160401b03811115611de657611de6614e8e565b604051908082528060200260200182016040528015611e0f578160200160208202803683370190505b50905060005b8551811015611ef957866001600160a01b03166304ec6351878381518110611e3f57611e3f615d4a565b602002602001015187868581518110611e5a57611e5a615d4a565b60200260200101516040518463ffffffff1660e01b8152600401611e8093929190615fd1565b602060405180830381865afa158015611e9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ec19190615fed565b6001600160c01b0316828281518110611edc57611edc615d4a565b602090810291909101015280611ef181615db1565b915050611e15565b5095945050505050565b611f0b614cc4565b600084611f685760405162461bcd60e51b815260206004820152603760248201526000805160206165b48339815191526044820152761c995cce88195b5c1d1e481c5d5bdc9d5b481a5b9c1d5d604a1b60648201526084016109fe565b60408301515185148015611f80575060a08301515185145b8015611f90575060c08301515185145b8015611fa0575060e08301515185145b61200a5760405162461bcd60e51b815260206004820152604160248201526000805160206165b483398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a4016109fe565b825151602084015151146120825760405162461bcd60e51b8152602060048201526044602482018190526000805160206165b4833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a4016109fe565b4363ffffffff168463ffffffff16106120f05760405162461bcd60e51b815260206004820152603c60248201526000805160206165b483398151915260448201527b7265733a20696e76616c6964207265666572656e636520626c6f636b60201b60648201526084016109fe565b604080518082019091526000808252602082015261210c614cc4565b866001600160401b0381111561212457612124614e8e565b60405190808252806020026020018201604052801561214d578160200160208202803683370190505b506020820152866001600160401b0381111561216b5761216b614e8e565b604051908082528060200260200182016040528015612194578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156121c8576121c8614e8e565b6040519080825280602002602001820160405280156121f1578160200160208202803683370190505b5081526020860151516001600160401b0381111561221157612211614e8e565b60405190808252806020026020018201604052801561223a578160200160208202803683370190505b508160200181905250600061230c8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa1580156122e3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612307919061609c565b614640565b905060005b876020015151811015612590576123568860200151828151811061233757612337615d4a565b6020026020010151805160009081526020918201519091526040902090565b8360200151828151811061236c5761236c615d4a565b6020908102919091010152801561242c57602083015161238d6001836160b9565b8151811061239d5761239d615d4a565b602002602001015160001c836020015182815181106123be576123be615d4a565b602002602001015160001c1161242c576040805162461bcd60e51b81526020600482015260248101919091526000805160206165b483398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016109fe565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec63518460200151838151811061247157612471615d4a565b60200260200101518b8b60000151858151811061249057612490615d4a565b60200260200101516040518463ffffffff1660e01b81526004016124b693929190615fd1565b602060405180830381865afa1580156124d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124f79190615fed565b6001600160c01b03168360000151828151811061251657612516615d4a565b60200260200101818152505061257c610daa612550848660000151858151811061254257612542615d4a565b6020026020010151166146c1565b8a60200151848151811061256657612566615d4a565b60200260200101516146ec90919063ffffffff16565b94508061258881615db1565b915050612311565b505061259b836147c4565b60975490935060ff166000816125b2576000612634565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612610573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126349190615d82565b905060005b8a811015612c7f578215612794578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f8681811061269057612690615d4a565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa1580156126d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126f49190615d82565b6126fe91906160d0565b116127945760405162461bcd60e51b815260206004820152606660248201526000805160206165b483398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c4016109fe565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106127d5576127d5615d4a565b9050013560f81c60f81b60f81c8c8c60a0015185815181106127f9576127f9615d4a565b60200260200101516040518463ffffffff1660e01b815260040161281f939291906160e8565b602060405180830381865afa15801561283c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128609190616109565b6001600160401b0319166128838a60400151838151811061233757612337615d4a565b6001600160401b0319161461291e5760405162461bcd60e51b815260206004820152606160248201526000805160206165b483398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016109fe565b61294e8960400151828151811061293757612937615d4a565b60200260200101518761418990919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061299157612991615d4a565b9050013560f81c60f81b60f81c8c8c60c0015185815181106129b5576129b5615d4a565b60200260200101516040518463ffffffff1660e01b81526004016129db939291906160e8565b602060405180830381865afa1580156129f8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a1c9190615e7b565b85602001518281518110612a3257612a32615d4a565b6001600160601b03909216602092830291909101820152850151805182908110612a5e57612a5e615d4a565b602002602001015185600001518281518110612a7c57612a7c615d4a565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612c6a57612af486600001518281518110612ac657612ac6615d4a565b60200260200101518f8f86818110612ae057612ae0615d4a565b600192013560f81c9290921c811614919050565b15612c58577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612b3a57612b3a615d4a565b9050013560f81c60f81b60f81c8e89602001518581518110612b5e57612b5e615d4a565b60200260200101518f60e001518881518110612b7c57612b7c615d4a565b60200260200101518781518110612b9557612b95615d4a565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612bf9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c1d9190615e7b565b8751805185908110612c3157612c31615d4a565b60200260200101818151612c459190616133565b6001600160601b03169052506001909101905b80612c6281615db1565b915050612aa0565b50508080612c7790615db1565b915050612639565b505050600080612c998c868a606001518b60800151610c8f565b9150915081612d0a5760405162461bcd60e51b815260206004820152604360248201526000805160206165b483398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016109fe565b80612d675760405162461bcd60e51b815260206004820152603960248201526000805160206165b48339815191526044820152781c995cce881cda59db985d1d5c99481a5cc81a5b9d985b1a59603a1b60648201526084016109fe565b50506000878260200151604051602001612d8292919061615b565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612db46145e1565b612dbe6000614008565b565b6000612dcf6020850185615061565b63ffffffff8116600090815260cb6020526040902054909150853590612e415760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b60648201526084016109fe565b8484604051602001612e54929190616285565b60408051601f19818403018152918152815160209283012063ffffffff8516600090815260cb90935291205414612ee15760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d6174636820746865604482015260008051602061655483398151915260648201526084016109fe565b63ffffffff8216600090815260cc602052604090205460ff1615612f795760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a4016109fe565b6064612f886020860186615061565b612f9291906162c3565b63ffffffff164363ffffffff16111561300d5760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527639b5903430b99030b63932b0b23c9032bc3834b932b21760491b60648201526084016109fe565b6001604051339063ffffffff8516907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a35050506132ed565b85518110156130975761306886828151811061233757612337615d4a565b82828151811061307a5761307a615d4a565b60209081029190910101528061308f81615db1565b91505061304a565b5060006130aa60408a0160208b01615061565b826040516020016130bc92919061615b565b604051602081830303815290604052805190602001209050866020013581146131665760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a4016109fe565b600086516001600160401b0381111561318157613181614e8e565b6040519080825280602002602001820160405280156131aa578160200160208202803683370190505b50905060005b875181101561329d577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae68583815181106131fa576131fa615d4a565b60200260200101516040518263ffffffff1660e01b815260040161322091815260200190565b602060405180830381865afa15801561323d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132619190615ca2565b82828151811061327357613273615d4a565b6001600160a01b03909216602092830291909101909101528061329581615db1565b9150506131b0565b5063ffffffff8616600081815260cc6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a35050505050505b50505050565b6033546001600160a01b031690565b60d0818154811061331257600080fd5b6000918252602090912001546001600160a01b0316905081565b60408051608081018252606081830181905260d354825263ffffffff438116602080850191909152908716918301919091528251601f85018290048202810182019093528383529091908490849081908401838280828437600092019190915250505050604080830191909152516133a89082906020016162eb565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907f76c838e3f0e8d8aae31520c434cc7eb96fa8d0e2e1828848fc42c82565d32fce9061340b9084906162eb565b60405180910390a260c9546134279063ffffffff1660016162c3565b60c9805463ffffffff191663ffffffff9290921691909117905550505050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061348257613482615d4a565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906134be908890869060040161607d565b600060405180830381865afa1580156134db573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526135039190810190615efa565b60008151811061351557613515615d4a565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015613581573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135a59190615fed565b6001600160c01b0316905060006135bb82614853565b9050816135c98a838a610f02565b9550955050505050935093915050565b606060d080548060200260200160405190810160405280929190818152602001828054801561363157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613613575b5050505050905090565b60cd546001600160a01b031633146136955760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064016109fe565b60006136a76040850160208601615061565b90503660006136b96040870187616376565b909250905060006136d06080880160608901615061565b905060ca60006136e36020890189615061565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161370f91906163bc565b60405160208183030381529060405280519060200120146137865760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d6174636820746865604482015260008051602061655483398151915260648201526084016109fe565b600060cb8161379860208a018a615061565b63ffffffff1663ffffffff16815260200190815260200160002054146138155760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b60648201526084016109fe565b61383f7f0000000000000000000000000000000000000000000000000000000000000000856162c3565b63ffffffff164363ffffffff1611156138b05760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b60648201526084016109fe565b6000866040516020016138c3919061645d565b6040516020818303038152906040528051906020012090506000806138eb8387878a8c611f03565b9150915060005b858110156139ea578460ff168360200151828151811061391457613914615d4a565b60200260200101516139269190616470565b6001600160601b031660648460000151838151811061394757613947615d4a565b60200260200101516001600160601b0316613962919061649f565b10156139d8576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d60648201526084016109fe565b806139e281615db1565b9150506138f2565b5060408051808201825263ffffffff43168152602080820184905291519091613a17918c918491016164be565b6040516020818303038152906040528051906020012060cb60008c6000016020810190613a449190615061565b63ffffffff1663ffffffff168152602001908152602001600020819055507f355781fa39bf42086618aba458ba14a142093e93eb7db6ed919b638cd35d1c018a82604051613a939291906164be565b60405180910390a15050600060d355505050505050505050565b600081815260d2602052604081205460ff1615613b005760405162461bcd60e51b815260206004820152601160248201527012185cda08185b1c9958591e481d5cd959607a1b60448201526064016109fe565b613b0a823461491f565b60d0805460018082019092557fe89d44c8fd6a9bac8af33ce47f56337617d449bf7ff3956b618c646de829cbcb0180546001600160a01b0319166001600160a01b038416908117909155600081815260d160209081526040808320805460ff19908116871790915588845260d2835292819020805490931690941790915582518681523491810191909152929350917f6c2525e4655984a225c978a74b5ad6bbe61d687428f2e5d73323956e139618a6910160405180910390a2919050565b613bd16145e1565b6001600160a01b038116613c365760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109fe565b610a1081614008565b613c476145e1565b610a108161405a565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613ca3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613cc79190615ca2565b6001600160a01b0316336001600160a01b031614613cf75760405162461bcd60e51b81526004016109fe90615cbf565b606654198119606654191614613d705760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d706044820152777420746f2070617573652066756e6374696f6e616c69747960401b60648201526084016109fe565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610b35565b613dd16040518060600160405280600063ffffffff16815260200160608152602001606081525090565b63ffffffff8216600090815260cb602052604090205480613e335760405162461bcd60e51b815260206004820152601c60248201527b15185cdac81c995cdc1bdb9cd948191bd95cc81b9bdd08195e1a5cdd60221b60448201526064016109fe565b50919050565b6001600160a01b038116613ec75760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016109fe565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b0316158015613f5157506001600160a01b03821615155b613fd35760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016109fe565b606681905560405181815233906000805160206165148339815191529060200160405180910390a261400482613e39565b5050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60cd80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f290600090a35050565b60ce80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f0ddfab8a635d71f15d72e2d2dff55d32119d13270d2ea4c3dc0043b66c2c476b90600090a35050565b614106614cde565b61410e614cf8565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561414157614143565bfe5b50806141815760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016109fe565b505092915050565b614191614cde565b614199614d16565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156141415750806141815760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016109fe565b614219614d34565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6142d9614cde565b5060408051808201909152600181526002602082015290565b6142fa614cde565b6000808061431660008051602061653483398151915286615d60565b90505b614322816149a0565b909350915060008051602061653483398151915282830983141561435c576040805180820190915290815260208101919091529392505050565b600080516020616534833981519152600182089050614319565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906143a8614d59565b60005b600281101561456d5760006143c182600661649f565b90508482600281106143d5576143d5615d4a565b602002015151836143e78360006160d0565b600c81106143f7576143f7615d4a565b602002015284826002811061440e5761440e615d4a565b6020020151602001518382600161442591906160d0565b600c811061443557614435615d4a565b602002015283826002811061444c5761444c615d4a565b602002015151518361445f8360026160d0565b600c811061446f5761446f615d4a565b602002015283826002811061448657614486615d4a565b602002015151600160200201518361449f8360036160d0565b600c81106144af576144af615d4a565b60200201528382600281106144c6576144c6615d4a565b6020020151602001516000600281106144e1576144e1615d4a565b6020020151836144f28360046160d0565b600c811061450257614502615d4a565b602002015283826002811061451957614519615d4a565b60200201516020015160016002811061453457614534615d4a565b6020020151836145458360056160d0565b600c811061455557614555615d4a565b6020020152508061456581615db1565b9150506143ab565b50614576614d78565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b336145ea6132f3565b6001600160a01b031614612dbe5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109fe565b60008061464c84614a22565b9050808360ff166001901b116146b85760405162461bcd60e51b815260206004820152603f602482015260008051602061659483398151915260448201527f69746d61703a206269746d61702065786365656473206d61782076616c75650060648201526084016109fe565b90505b92915050565b6000805b82156146bb576146d66001846160b9565b90921691806146e4816164f1565b9150506146c5565b6146f4614cde565b6102008261ffff161061473c5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016109fe565b8161ffff16600114156147505750816146bb565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff16106147b957600161ffff871660ff83161c8116141561479c576147998484614189565b93505b6147a68384614189565b92506201fffe600192831b16910161476c565b509195945050505050565b6147cc614cde565b81511580156147dd57506020820151155b156147fb575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020616534833981519152846020015161482e9190615d60565b614846906000805160206165348339815191526160b9565b905292915050565b919050565b6060600080614861846146c1565b61ffff166001600160401b0381111561487c5761487c614e8e565b6040519080825280601f01601f1916602001820160405280156148a6576020820181803683370190505b5090506000805b8251821080156148be575061010081105b15614915576001811b935085841615614905578060f81b8383815181106148e7576148e7615d4a565b60200101906001600160f81b031916908160001a9053508160010191505b61490e81615db1565b90506148ad565b5090949350505050565b60cf54600090839061493a906001600160a01b031682614b8b565b604051639498bd7160e01b8152600481018690529092506001600160a01b03831690639498bd719085906024016000604051808303818588803b15801561498057600080fd5b505af1158015614994573d6000803e3d6000fd5b50505050505092915050565b60008080600080516020616534833981519152600360008051602061653483398151915286600080516020616534833981519152888909090890506000614a16827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020616534833981519152614ba0565b91959194509092505050565b600061010082511115614a995760405162461bcd60e51b815260206004820152604460248201819052600080516020616594833981519152908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a4016109fe565b8151614aa757506000919050565b60008083600081518110614abd57614abd615d4a565b0160200151600160f89190911c81901b92505b8451811015614b8257848181518110614aeb57614aeb615d4a565b0160200151600160f89190911c1b9150828211614b6e5760405162461bcd60e51b8152602060048201526047602482015260008051602061659483398151915260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a4016109fe565b91811791614b7b81615db1565b9050614ad0565b50909392505050565b6000614b9960008484614c45565b9392505050565b600080614bab614d78565b614bb3614d96565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015614141575082614c3a5760405162461bcd60e51b815260206004820152601a602482015279424e3235342e6578704d6f643a2063616c6c206661696c75726560301b60448201526064016109fe565b505195945050505050565b60006c5af43d3d93803e602a57fd5bf36021528260145273602c3d8160093d39f33d3d3d3d363d3d37363d73600052816035600c86f5905080614c905763301164256000526004601cfd5b60006021529392505050565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b604051806040016040528060608152602001606081525090565b604051806040016040528060008152602001600081525090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280614d47614db4565b8152602001614d54614db4565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b0381168114610a1057600080fd5b600060208284031215614df957600080fd5b81356146b881614dd2565b600060208284031215614e1657600080fd5b5035919050565b600080600080600060a08688031215614e3557600080fd5b8535614e4081614dd2565b94506020860135614e5081614dd2565b93506040860135614e6081614dd2565b92506060860135614e7081614dd2565b91506080860135614e8081614dd2565b809150509295509295909350565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614ec657614ec6614e8e565b60405290565b60405161010081016001600160401b0381118282101715614ec657614ec6614e8e565b604051601f8201601f191681016001600160401b0381118282101715614f1757614f17614e8e565b604052919050565b600060408284031215614f3157600080fd5b614f39614ea4565b9050813581526020820135602082015292915050565b600082601f830112614f6057600080fd5b604080519081016001600160401b0381118282101715614f8257614f82614e8e565b8060405250806040840185811115614f9957600080fd5b845b818110156147b9578035835260209283019201614f9b565b600060808284031215614fc557600080fd5b614fcd614ea4565b9050614fd98383614f4f565b8152614fe88360408401614f4f565b602082015292915050565b600080600080610120858703121561500a57600080fd5b8435935061501b8660208701614f1f565b925061502a8660608701614fb3565b91506150398660e08701614f1f565b905092959194509250565b63ffffffff81168114610a1057600080fd5b803561484e81615044565b60006020828403121561507357600080fd5b81356146b881615044565b6001600160a01b0391909116815260200190565b60006001600160401b038211156150ab576150ab614e8e565b5060051b60200190565b600080604083850312156150c857600080fd5b82356150d381614dd2565b91506020838101356001600160401b038111156150ef57600080fd5b8401601f8101861361510057600080fd5b803561511361510e82615092565b614eef565b81815260059190911b8201830190838101908883111561513257600080fd5b928401925b8284101561515957833561514a81614dd2565b82529284019290840190615137565b80955050505050509250929050565b600081518084526020808501945080840160005b838110156151985781518752958201959082019060010161517c565b509495945050505050565b602081526000614b996020830184615168565b6000806000606084860312156151cb57600080fd5b83356151d681614dd2565b92506020848101356001600160401b03808211156151f357600080fd5b818701915087601f83011261520757600080fd5b81358181111561521957615219614e8e565b61522b601f8201601f19168501614eef565b9150808252888482850101111561524157600080fd5b808484018584013760008482840101525080945050505061526460408501615056565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015615303578385038a52825180518087529087019087870190845b818110156152ee57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b031690840152928901926060909201916001016152aa565b50509a87019a9550509185019160010161528c565b509298975050505050505050565b602081526000614b99602083018461526d565b8015158114610a1057600080fd5b60006020828403121561534457600080fd5b81356146b881615324565b600082601f83011261536057600080fd5b8135602061537061510e83615092565b82815260059290921b8401810191818101908684111561538f57600080fd5b8286015b848110156153aa5780358352918301918301615393565b509695505050505050565b600080604083850312156153c857600080fd5b82356153d381614dd2565b915060208301356001600160401b038111156153ee57600080fd5b6153fa8582860161534f565b9150509250929050565b600081518084526020808501945080840160005b838110156151985781516001600160a01b031687529582019590820190600101615418565b602081526000614b996020830184615404565b60008083601f84011261546257600080fd5b5081356001600160401b0381111561547957600080fd5b60208301915083602082850101111561549157600080fd5b9250929050565b600080600080600080608087890312156154b157600080fd5b86356154bc81614dd2565b955060208701356154cc81615044565b945060408701356001600160401b03808211156154e857600080fd5b6154f48a838b01615450565b9096509450606089013591508082111561550d57600080fd5b818901915089601f83011261552157600080fd5b81358181111561553057600080fd5b8a60208260051b850101111561554557600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b8381101561519857815163ffffffff168752958201959082019060010161556f565b6000602080835283516080828501526155ad60a085018261555b565b905081850151601f19808684030160408701526155ca838361555b565b925060408701519150808684030160608701526155e7838361555b565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b8281101561563e578487830301845261562c82875161555b565b95880195938801939150600101615612565b509998505050505050505050565b60ff81168114610a1057600080fd5b60006020828403121561566d57600080fd5b81356146b88161564c565b60008060006060848603121561568d57600080fd5b833561569881614dd2565b925060208401356001600160401b038111156156b357600080fd5b6156bf8682870161534f565b92505060408401356156d081615044565b809150509250925092565b600082601f8301126156ec57600080fd5b813560206156fc61510e83615092565b82815260059290921b8401810191818101908684111561571b57600080fd5b8286015b848110156153aa57803561573281615044565b835291830191830161571f565b600082601f83011261575057600080fd5b8135602061576061510e83615092565b82815260069290921b8401810191818101908684111561577f57600080fd5b8286015b848110156153aa576157958882614f1f565b835291830191604001615783565b600082601f8301126157b457600080fd5b813560206157c461510e83615092565b82815260059290921b840181019181810190868411156157e357600080fd5b8286015b848110156153aa5780356001600160401b038111156158065760008081fd5b6158148986838b01016156db565b8452509183019183016157e7565b6000610180828403121561583557600080fd5b61583d614ecc565b905081356001600160401b038082111561585657600080fd5b615862858386016156db565b8352602084013591508082111561587857600080fd5b6158848583860161573f565b6020840152604084013591508082111561589d57600080fd5b6158a98583860161573f565b60408401526158bb8560608601614fb3565b60608401526158cd8560e08601614f1f565b60808401526101208401359150808211156158e757600080fd5b6158f3858386016156db565b60a084015261014084013591508082111561590d57600080fd5b615919858386016156db565b60c084015261016084013591508082111561593357600080fd5b50615940848285016157a3565b60e08301525092915050565b60008060008060006080868803121561596457600080fd5b8535945060208601356001600160401b038082111561598257600080fd5b61598e89838a01615450565b9096509450604088013591506159a382615044565b909250606087013590808211156159b957600080fd5b506159c688828901615822565b9150509295509295909350565b600081518084526020808501945080840160005b838110156151985781516001600160601b0316875295820195908201906001016159e7565b6040815260008351604080840152615a2760808401826159d3565b90506020850151603f19848303016060850152615a4482826159d3565b925050508260208301529392505050565b600060808284031215613e3357600080fd5b600060608284031215613e3357600080fd5b60008060008084860360a0811215615a9057600080fd5b85356001600160401b0380821115615aa757600080fd5b615ab389838a01615a55565b96506020880135915080821115615ac957600080fd5b615ad589838a01615a67565b95506040603f1984011215615ae957600080fd5b6040880194506080880135925080831115615b0357600080fd5b5050615b118782880161573f565b91505092959194509250565b600080600060408486031215615b3257600080fd5b8335615b3d81615044565b925060208401356001600160401b03811115615b5857600080fd5b615b6486828701615450565b9497909650939450505050565b600080600060608486031215615b8657600080fd5b8335615b9181614dd2565b92506020840135915060408401356156d081615044565b828152604060208201526000615bc1604083018461526d565b949350505050565b600080600060608486031215615bde57600080fd5b83356001600160401b0380821115615bf557600080fd5b615c0187838801615a55565b94506020860135915080821115615c1757600080fd5b615c2387838801615a67565b93506040860135915080821115615c3957600080fd5b50615c4686828701615822565b9150509250925092565b6020815263ffffffff82511660208201526000602083015160606040840152615c7c6080840182615404565b90506040840151601f19848303016060850152615c998282615168565b95945050505050565b600060208284031215615cb457600080fd5b81516146b881614dd2565b6020808252602a9082015260008051602061657483398151915260408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215615d0957600080fd5b81516146b881615324565b602080825260289082015260008051602061657483398151915260408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600082615d7d57634e487b7160e01b600052601260045260246000fd5b500690565b600060208284031215615d9457600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b6000600019821415615dc557615dc5615d9b565b5060010190565b60006020808385031215615ddf57600080fd5b82516001600160401b03811115615df557600080fd5b8301601f81018513615e0657600080fd5b8051615e1461510e82615092565b81815260059190911b82018301908381019087831115615e3357600080fd5b928401925b82841015615e5157835182529284019290840190615e38565b979650505050505050565b92835260ff91909116602083015263ffffffff16604082015260600190565b600060208284031215615e8d57600080fd5b81516001600160601b03811681146146b857600080fd5b81835260006001600160fb1b03831115615ebd57600080fd5b8260051b8083602087013760009401602001938452509192915050565b63ffffffff84168152604060208201526000615c99604083018486615ea4565b60006020808385031215615f0d57600080fd5b82516001600160401b03811115615f2357600080fd5b8301601f81018513615f3457600080fd5b8051615f4261510e82615092565b81815260059190911b82018301908381019087831115615f6157600080fd5b928401925b82841015615e51578351615f7981615044565b82529284019290840190615f66565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000615c99604083018486615f88565b92835263ffffffff918216602084015216604082015260600190565b600060208284031215615fff57600080fd5b81516001600160c01b03811681146146b857600080fd5b60006020828403121561602857600080fd5b81516146b881615044565b600060ff821660ff81141561604a5761604a615d9b565b60010192915050565b604081526000616067604083018587615f88565b905063ffffffff83166020830152949350505050565b63ffffffff83168152604060208201526000615bc16040830184615168565b6000602082840312156160ae57600080fd5b81516146b88161564c565b6000828210156160cb576160cb615d9b565b500390565b600082198211156160e3576160e3615d9b565b500190565b60ff93909316835263ffffffff918216602084015216604082015260600190565b60006020828403121561611b57600080fd5b81516001600160401b0319811681146146b857600080fd5b60006001600160601b038381169083168181101561615357616153615d9b565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156161965781518552938201939082019060010161617a565b5092979650505050505050565b6000808335601e198436030181126161ba57600080fd5b83016020810192503590506001600160401b038111156161d957600080fd5b8060051b360383131561549157600080fd5b60006060830182356161fc81615044565b63ffffffff1684526020616212848201856161a3565b6060878401529283905291600090608087015b8183101561625557843561623881614dd2565b6001600160a01b0316815293830193600192909201918301616225565b61626260408801886161a3565b955093508781036040890152616279818686615ea4565b98975050505050505050565b60608152600061629860608301856161eb565b905082356162a581615044565b63ffffffff8116602084015250602083013560408301529392505050565b600063ffffffff8083168185168083038211156162e2576162e2615d9b565b01949350505050565b6000602080835283518184015263ffffffff8185015116604084015260408401516080606085015280518060a086015260005b8181101561633a5782810184015186820160c00152830161631e565b8181111561634c57600060c083880101525b50606086015163ffffffff811660808701529250601f01601f19169390930160c001949350505050565b6000808335601e1984360301811261638d57600080fd5b8301803591506001600160401b038211156163a757600080fd5b60200191503681900382131561549157600080fd5b6020815281356020820152600060208301356163d781615044565b63ffffffff81166040840152506040830135601e198436030181126163fb57600080fd5b830180356001600160401b0381111561641357600080fd5b80360385131561642257600080fd5b6080606085015261643a60a085018260208501615f88565b91505061644960608501615056565b63ffffffff81166080850152509392505050565b602081526000614b9960208301846161eb565b60006001600160601b038281168482168115158284048211161561649657616496615d9b565b02949350505050565b60008160001904831182151516156164b9576164b9615d9b565b500290565b6060815260006164d160608301856161eb565b905063ffffffff8351166020830152602083015160408301529392505050565b600061ffff8083168181141561650957616509615d9b565b600101939250505056feab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47206f6e65207265636f7264656420696e2074686520636f6e74726163740000006d73672e73656e646572206973206e6f74207065726d697373696f6e656420614269746d61705574696c732e6f72646572656442797465734172726179546f42424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220f83d103d693e4a30f92b9201c2b0e8958a993bcae35e350ee8bada27a81a910364736f6c634300080c0033",
}

// ContractIncredibleSquaringTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIncredibleSquaringTaskManagerMetaData.ABI instead.
var ContractIncredibleSquaringTaskManagerABI = ContractIncredibleSquaringTaskManagerMetaData.ABI

// ContractIncredibleSquaringTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIncredibleSquaringTaskManagerMetaData.Bin instead.
var ContractIncredibleSquaringTaskManagerBin = ContractIncredibleSquaringTaskManagerMetaData.Bin

// DeployContractIncredibleSquaringTaskManager deploys a new Ethereum contract, binding an instance of ContractIncredibleSquaringTaskManager to it.
func DeployContractIncredibleSquaringTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractIncredibleSquaringTaskManager, error) {
	parsed, err := ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIncredibleSquaringTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIncredibleSquaringTaskManager{ContractIncredibleSquaringTaskManagerCaller: ContractIncredibleSquaringTaskManagerCaller{contract: contract}, ContractIncredibleSquaringTaskManagerTransactor: ContractIncredibleSquaringTaskManagerTransactor{contract: contract}, ContractIncredibleSquaringTaskManagerFilterer: ContractIncredibleSquaringTaskManagerFilterer{contract: contract}}, nil
}

// ContractIncredibleSquaringTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManager struct {
	ContractIncredibleSquaringTaskManagerCaller     // Read-only binding to the contract
	ContractIncredibleSquaringTaskManagerTransactor // Write-only binding to the contract
	ContractIncredibleSquaringTaskManagerFilterer   // Log filterer for contract events
}

// ContractIncredibleSquaringTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIncredibleSquaringTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIncredibleSquaringTaskManagerSession struct {
	Contract     *ContractIncredibleSquaringTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                          // Call options to use throughout this session
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractIncredibleSquaringTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIncredibleSquaringTaskManagerCallerSession struct {
	Contract *ContractIncredibleSquaringTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                // Call options to use throughout this session
}

// ContractIncredibleSquaringTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIncredibleSquaringTaskManagerTransactorSession struct {
	Contract     *ContractIncredibleSquaringTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                // Transaction auth options to use throughout this session
}

// ContractIncredibleSquaringTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerRaw struct {
	Contract *ContractIncredibleSquaringTaskManager // Generic contract binding to access the raw methods on
}

// ContractIncredibleSquaringTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerCallerRaw struct {
	Contract *ContractIncredibleSquaringTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIncredibleSquaringTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerTransactorRaw struct {
	Contract *ContractIncredibleSquaringTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIncredibleSquaringTaskManager creates a new instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManager(address common.Address, backend bind.ContractBackend) (*ContractIncredibleSquaringTaskManager, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManager{ContractIncredibleSquaringTaskManagerCaller: ContractIncredibleSquaringTaskManagerCaller{contract: contract}, ContractIncredibleSquaringTaskManagerTransactor: ContractIncredibleSquaringTaskManagerTransactor{contract: contract}, ContractIncredibleSquaringTaskManagerFilterer: ContractIncredibleSquaringTaskManagerFilterer{contract: contract}}, nil
}

// NewContractIncredibleSquaringTaskManagerCaller creates a new read-only instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIncredibleSquaringTaskManagerCaller, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerCaller{contract: contract}, nil
}

// NewContractIncredibleSquaringTaskManagerTransactor creates a new write-only instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIncredibleSquaringTaskManagerTransactor, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTransactor{contract: contract}, nil
}

// NewContractIncredibleSquaringTaskManagerFilterer creates a new log filterer instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIncredibleSquaringTaskManagerFilterer, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerFilterer{contract: contract}, nil
}

// bindContractIncredibleSquaringTaskManager binds a generic wrapper to an already deployed contract.
func bindContractIncredibleSquaringTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSquaringTaskManager.Contract.ContractIncredibleSquaringTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.ContractIncredibleSquaringTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.ContractIncredibleSquaringTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSquaringTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")
	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")
	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "aggregator")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Aggregator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Aggregator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)
	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskHashes(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskHashes(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)
	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskResponses(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskResponses(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "blsApkRegistry")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.BlsApkRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.BlsApkRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)
	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CheckSignatures(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CheckSignatures(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "delegation")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Delegation(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Delegation(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "generator")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Generator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Generator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Generator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() view returns(address[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetAllPools(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getAllPools")
	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() view returns(address[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetAllPools() ([]common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetAllPools(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() view returns(address[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetAllPools() ([]common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetAllPools(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)
	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetBatchOperatorFromId(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetBatchOperatorFromId(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)
	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetBatchOperatorId(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetBatchOperatorId(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)
	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)
	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)
	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponse is a free data retrieval call binding the contract method 0xffb3959f.
//
// Solidity: function getTaskResponse(uint32 taskIndex) view returns((uint32,address[],uint256[]) taskResponse)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetTaskResponse(opts *bind.CallOpts, taskIndex uint32) (IIncredibleSquaringTaskManagerTaskResponse, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getTaskResponse", taskIndex)
	if err != nil {
		return *new(IIncredibleSquaringTaskManagerTaskResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(IIncredibleSquaringTaskManagerTaskResponse)).(*IIncredibleSquaringTaskManagerTaskResponse)

	return out0, err
}

// GetTaskResponse is a free data retrieval call binding the contract method 0xffb3959f.
//
// Solidity: function getTaskResponse(uint32 taskIndex) view returns((uint32,address[],uint256[]) taskResponse)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetTaskResponse(taskIndex uint32) (IIncredibleSquaringTaskManagerTaskResponse, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetTaskResponse(&_ContractIncredibleSquaringTaskManager.CallOpts, taskIndex)
}

// GetTaskResponse is a free data retrieval call binding the contract method 0xffb3959f.
//
// Solidity: function getTaskResponse(uint32 taskIndex) view returns((uint32,address[],uint256[]) taskResponse)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetTaskResponse(taskIndex uint32) (IIncredibleSquaringTaskManagerTaskResponse, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetTaskResponse(&_ContractIncredibleSquaringTaskManager.CallOpts, taskIndex)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")
	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) IsPool(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "isPool", arg0)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) IsPool(arg0 common.Address) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.IsPool(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) IsPool(arg0 common.Address) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.IsPool(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "latestTaskNum")
	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.LatestTaskNum(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.LatestTaskNum(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "owner")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Owner(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Owner(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "paused", index)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused(&_ContractIncredibleSquaringTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused(&_ContractIncredibleSquaringTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "paused0")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused0(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused0(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "pauserRegistry")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauserRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauserRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// PoolImplementation is a free data retrieval call binding the contract method 0xcefa7799.
//
// Solidity: function poolImplementation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) PoolImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "poolImplementation")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// PoolImplementation is a free data retrieval call binding the contract method 0xcefa7799.
//
// Solidity: function poolImplementation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) PoolImplementation() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PoolImplementation(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// PoolImplementation is a free data retrieval call binding the contract method 0xcefa7799.
//
// Solidity: function poolImplementation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) PoolImplementation() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PoolImplementation(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "pools", arg0)
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Pools(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Pools(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "registryCoordinator")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "stakeRegistry")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StakeRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StakeRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "staleStakesForbidden")
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "taskNumber")
	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskNumber(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskNumber(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TaskSuccessfullyChallenged is a free data retrieval call binding the contract method 0x1ce8a3e3.
//
// Solidity: function taskSuccessfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TaskSuccessfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "taskSuccessfullyChallenged", arg0)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// TaskSuccessfullyChallenged is a free data retrieval call binding the contract method 0x1ce8a3e3.
//
// Solidity: function taskSuccessfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TaskSuccessfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskSuccessfullyChallenged(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// TaskSuccessfullyChallenged is a free data retrieval call binding the contract method 0x1ce8a3e3.
//
// Solidity: function taskSuccessfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TaskSuccessfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskSuccessfullyChallenged(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error,
) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error,
) {
	return _ContractIncredibleSquaringTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error,
) {
	return _ContractIncredibleSquaringTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) UsedHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "usedHashes", arg0)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) UsedHashes(arg0 [32]byte) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.UsedHashes(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) UsedHashes(arg0 [32]byte) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.UsedHashes(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xbb972b04.
//
// Solidity: function createNewTask(uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "createNewTask", quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xbb972b04.
//
// Solidity: function createNewTask(uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) CreateNewTask(quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CreateNewTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xbb972b04.
//
// Solidity: function createNewTask(uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) CreateNewTask(quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CreateNewTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, quorumThresholdPercentage, quorumNumbers)
}

// CreatePool is a paid mutator transaction binding the contract method 0xee2373e9.
//
// Solidity: function createPool(bytes32 hash) payable returns(address poolAddress)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) CreatePool(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "createPool", hash)
}

// CreatePool is a paid mutator transaction binding the contract method 0xee2373e9.
//
// Solidity: function createPool(bytes32 hash) payable returns(address poolAddress)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) CreatePool(hash [32]byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CreatePool(&_ContractIncredibleSquaringTaskManager.TransactOpts, hash)
}

// CreatePool is a paid mutator transaction binding the contract method 0xee2373e9.
//
// Solidity: function createPool(bytes32 hash) payable returns(address poolAddress)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) CreatePool(hash [32]byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CreatePool(&_ContractIncredibleSquaringTaskManager.TransactOpts, hash)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _poolImplementation) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _poolImplementation common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator, _poolImplementation)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _poolImplementation) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _poolImplementation common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Initialize(&_ContractIncredibleSquaringTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator, _poolImplementation)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator, address _poolImplementation) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address, _poolImplementation common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Initialize(&_ContractIncredibleSquaringTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator, _poolImplementation)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Pause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Pause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauseAll(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauseAll(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// QueueFileHash is a paid mutator transaction binding the contract method 0x8889df3c.
//
// Solidity: function queueFileHash(bytes32 fileHash) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) QueueFileHash(opts *bind.TransactOpts, fileHash [32]byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "queueFileHash", fileHash)
}

// QueueFileHash is a paid mutator transaction binding the contract method 0x8889df3c.
//
// Solidity: function queueFileHash(bytes32 fileHash) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) QueueFileHash(fileHash [32]byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.QueueFileHash(&_ContractIncredibleSquaringTaskManager.TransactOpts, fileHash)
}

// QueueFileHash is a paid mutator transaction binding the contract method 0x8889df3c.
//
// Solidity: function queueFileHash(bytes32 fileHash) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) QueueFileHash(fileHash [32]byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.QueueFileHash(&_ContractIncredibleSquaringTaskManager.TransactOpts, fileHash)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x8866d187.
//
// Solidity: function raiseAndResolveChallenge((bytes32,uint32,bytes,uint32) task, (uint32,address[],uint256[]) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, taskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x8866d187.
//
// Solidity: function raiseAndResolveChallenge((bytes32,uint32,bytes,uint32) task, (uint32,address[],uint256[]) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RaiseAndResolveChallenge(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, taskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x8866d187.
//
// Solidity: function raiseAndResolveChallenge((bytes32,uint32,bytes,uint32) task, (uint32,address[],uint256[]) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) RaiseAndResolveChallenge(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, taskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RenounceOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RenounceOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xe25a54a9.
//
// Solidity: function respondToTask((bytes32,uint32,bytes,uint32) task, (uint32,address[],uint256[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xe25a54a9.
//
// Solidity: function respondToTask((bytes32,uint32,bytes,uint32) task, (uint32,address[],uint256[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RespondToTask(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RespondToTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xe25a54a9.
//
// Solidity: function respondToTask((bytes32,uint32,bytes,uint32) task, (uint32,address[],uint256[]) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) RespondToTask(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RespondToTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address newAggregator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) SetAggregator(opts *bind.TransactOpts, newAggregator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "setAggregator", newAggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address newAggregator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) SetAggregator(newAggregator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetAggregator(&_ContractIncredibleSquaringTaskManager.TransactOpts, newAggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address newAggregator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) SetAggregator(newAggregator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetAggregator(&_ContractIncredibleSquaringTaskManager.TransactOpts, newAggregator)
}

// SetGenerator is a paid mutator transaction binding the contract method 0x4a7c7e4b.
//
// Solidity: function setGenerator(address newGenerator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) SetGenerator(opts *bind.TransactOpts, newGenerator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "setGenerator", newGenerator)
}

// SetGenerator is a paid mutator transaction binding the contract method 0x4a7c7e4b.
//
// Solidity: function setGenerator(address newGenerator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) SetGenerator(newGenerator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetGenerator(&_ContractIncredibleSquaringTaskManager.TransactOpts, newGenerator)
}

// SetGenerator is a paid mutator transaction binding the contract method 0x4a7c7e4b.
//
// Solidity: function setGenerator(address newGenerator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) SetGenerator(newGenerator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetGenerator(&_ContractIncredibleSquaringTaskManager.TransactOpts, newGenerator)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetPauserRegistry(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetPauserRegistry(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetStaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetStaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TransferOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TransferOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Unpause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Unpause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator is returned from FilterAggregatorUpdated and is used to iterate over the raw logs and unpacked data for AggregatorUpdated events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerAggregatorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerAggregatorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerAggregatorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerAggregatorUpdated represents a AggregatorUpdated event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerAggregatorUpdated struct {
	OldAggregator common.Address
	NewAggregator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAggregatorUpdated is a free log retrieval operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterAggregatorUpdated(opts *bind.FilterOpts, oldAggregator []common.Address, newAggregator []common.Address) (*ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator, error) {
	var oldAggregatorRule []interface{}
	for _, oldAggregatorItem := range oldAggregator {
		oldAggregatorRule = append(oldAggregatorRule, oldAggregatorItem)
	}
	var newAggregatorRule []interface{}
	for _, newAggregatorItem := range newAggregator {
		newAggregatorRule = append(newAggregatorRule, newAggregatorItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "AggregatorUpdated", oldAggregatorRule, newAggregatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "AggregatorUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregatorUpdated is a free log subscription operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerAggregatorUpdated, oldAggregator []common.Address, newAggregator []common.Address) (event.Subscription, error) {
	var oldAggregatorRule []interface{}
	for _, oldAggregatorItem := range oldAggregator {
		oldAggregatorRule = append(oldAggregatorRule, oldAggregatorItem)
	}
	var newAggregatorRule []interface{}
	for _, newAggregatorItem := range newAggregator {
		newAggregatorRule = append(newAggregatorRule, newAggregatorItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "AggregatorUpdated", oldAggregatorRule, newAggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerAggregatorUpdated)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAggregatorUpdated is a log parse operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseAggregatorUpdated(log types.Log) (*ContractIncredibleSquaringTaskManagerAggregatorUpdated, error) {
	event := new(ContractIncredibleSquaringTaskManagerAggregatorUpdated)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator is returned from FilterGeneratorUpdated and is used to iterate over the raw logs and unpacked data for GeneratorUpdated events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerGeneratorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerGeneratorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerGeneratorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerGeneratorUpdated represents a GeneratorUpdated event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerGeneratorUpdated struct {
	OldGenerator common.Address
	NewGenerator common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGeneratorUpdated is a free log retrieval operation binding the contract event 0x0ddfab8a635d71f15d72e2d2dff55d32119d13270d2ea4c3dc0043b66c2c476b.
//
// Solidity: event GeneratorUpdated(address indexed oldGenerator, address indexed newGenerator)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterGeneratorUpdated(opts *bind.FilterOpts, oldGenerator []common.Address, newGenerator []common.Address) (*ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator, error) {
	var oldGeneratorRule []interface{}
	for _, oldGeneratorItem := range oldGenerator {
		oldGeneratorRule = append(oldGeneratorRule, oldGeneratorItem)
	}
	var newGeneratorRule []interface{}
	for _, newGeneratorItem := range newGenerator {
		newGeneratorRule = append(newGeneratorRule, newGeneratorItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "GeneratorUpdated", oldGeneratorRule, newGeneratorRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "GeneratorUpdated", logs: logs, sub: sub}, nil
}

// WatchGeneratorUpdated is a free log subscription operation binding the contract event 0x0ddfab8a635d71f15d72e2d2dff55d32119d13270d2ea4c3dc0043b66c2c476b.
//
// Solidity: event GeneratorUpdated(address indexed oldGenerator, address indexed newGenerator)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchGeneratorUpdated(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerGeneratorUpdated, oldGenerator []common.Address, newGenerator []common.Address) (event.Subscription, error) {
	var oldGeneratorRule []interface{}
	for _, oldGeneratorItem := range oldGenerator {
		oldGeneratorRule = append(oldGeneratorRule, oldGeneratorItem)
	}
	var newGeneratorRule []interface{}
	for _, newGeneratorItem := range newGenerator {
		newGeneratorRule = append(newGeneratorRule, newGeneratorItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "GeneratorUpdated", oldGeneratorRule, newGeneratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerGeneratorUpdated)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "GeneratorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGeneratorUpdated is a log parse operation binding the contract event 0x0ddfab8a635d71f15d72e2d2dff55d32119d13270d2ea4c3dc0043b66c2c476b.
//
// Solidity: event GeneratorUpdated(address indexed oldGenerator, address indexed newGenerator)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseGeneratorUpdated(log types.Log) (*ContractIncredibleSquaringTaskManagerGeneratorUpdated, error) {
	event := new(ContractIncredibleSquaringTaskManagerGeneratorUpdated)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "GeneratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerInitializedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerInitialized represents a Initialized event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerInitializedIterator, error) {
	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerInitializedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerInitialized) (event.Subscription, error) {
	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerInitialized)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractIncredibleSquaringTaskManagerInitialized, error) {
	event := new(ContractIncredibleSquaringTaskManagerInitialized)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IIncredibleSquaringTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x76c838e3f0e8d8aae31520c434cc7eb96fa8d0e2e1828848fc42c82565d32fce.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (bytes32,uint32,bytes,uint32) task)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator, error) {
	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x76c838e3f0e8d8aae31520c434cc7eb96fa8d0e2e1828848fc42c82565d32fce.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (bytes32,uint32,bytes,uint32) task)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {
	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTaskCreated is a log parse operation binding the contract event 0x76c838e3f0e8d8aae31520c434cc7eb96fa8d0e2e1828848fc42c82565d32fce.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (bytes32,uint32,bytes,uint32) task)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractIncredibleSquaringTaskManagerNewTaskCreated, error) {
	event := new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator struct {
	Event *ContractIncredibleSquaringTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractIncredibleSquaringTaskManagerOwnershipTransferred, error) {
	event := new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPausedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerPaused represents a Paused event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleSquaringTaskManagerPausedIterator, error) {
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerPausedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerPaused, account []common.Address) (event.Subscription, error) {
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerPaused)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParsePaused(log types.Log) (*ContractIncredibleSquaringTaskManagerPaused, error) {
	event := new(ContractIncredibleSquaringTaskManagerPaused)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator struct {
	Event *ContractIncredibleSquaringTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator, error) {
	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerPauserRegistrySet) (event.Subscription, error) {
	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractIncredibleSquaringTaskManagerPauserRegistrySet, error) {
	event := new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPoolCreatedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerPoolCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerPoolCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerPoolCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerPoolCreated represents a PoolCreated event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPoolCreated struct {
	PoolAddress common.Address
	Hash        [32]byte
	Balance     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x6c2525e4655984a225c978a74b5ad6bbe61d687428f2e5d73323956e139618a6.
//
// Solidity: event PoolCreated(address indexed poolAddress, bytes32 hash, uint256 balance)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterPoolCreated(opts *bind.FilterOpts, poolAddress []common.Address) (*ContractIncredibleSquaringTaskManagerPoolCreatedIterator, error) {
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "PoolCreated", poolAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerPoolCreatedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x6c2525e4655984a225c978a74b5ad6bbe61d687428f2e5d73323956e139618a6.
//
// Solidity: event PoolCreated(address indexed poolAddress, bytes32 hash, uint256 balance)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerPoolCreated, poolAddress []common.Address) (event.Subscription, error) {
	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "PoolCreated", poolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerPoolCreated)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "PoolCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolCreated is a log parse operation binding the contract event 0x6c2525e4655984a225c978a74b5ad6bbe61d687428f2e5d73323956e139618a6.
//
// Solidity: event PoolCreated(address indexed poolAddress, bytes32 hash, uint256 balance)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParsePoolCreated(log types.Log) (*ContractIncredibleSquaringTaskManagerPoolCreated, error) {
	event := new(ContractIncredibleSquaringTaskManagerPoolCreated)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator, error) {
	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {
	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator, error) {
	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {
	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator, error) {
	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {
	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskCompletedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSquaringTaskManagerTaskCompletedIterator, error) {
	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskCompletedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {
	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskCompleted)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskCompleted, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskCompleted)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskRespondedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskResponded represents a TaskResponded event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskResponded struct {
	TaskResponse         IIncredibleSquaringTaskManagerTaskResponse
	TaskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x355781fa39bf42086618aba458ba14a142093e93eb7db6ed919b638cd35d1c01.
//
// Solidity: event TaskResponded((uint32,address[],uint256[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerTaskRespondedIterator, error) {
	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskRespondedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x355781fa39bf42086618aba458ba14a142093e93eb7db6ed919b638cd35d1c01.
//
// Solidity: event TaskResponded((uint32,address[],uint256[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskResponded) (event.Subscription, error) {
	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskResponded)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskResponded is a log parse operation binding the contract event 0x355781fa39bf42086618aba458ba14a142093e93eb7db6ed919b638cd35d1c01.
//
// Solidity: event TaskResponded((uint32,address[],uint256[]) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskResponded, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskResponded)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerUnpausedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSquaringTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSquaringTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSquaringTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerUnpaused represents a Unpaused event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleSquaringTaskManagerUnpausedIterator, error) {
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerUnpausedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerUnpaused)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractIncredibleSquaringTaskManagerUnpaused, error) {
	event := new(ContractIncredibleSquaringTaskManagerUnpaused)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
