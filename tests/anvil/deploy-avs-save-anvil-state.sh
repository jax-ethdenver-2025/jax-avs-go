#!/bin/bash

RPC_URL=http://localhost:8545
PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# cd to the directory of this script so that this can be run from anywhere
parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)
cd "$parent_path"

set -a
source ./utils.sh
set +a

# cleanup() {
#     echo "Executing cleanup function..."
#     set +e
#     docker rm -f anvil
#     exit_status=$?
#     if [ $exit_status -ne 0 ]; then
#         echo "Script exited due to set -e on line $1 with command '$2'. Exit status: $exit_status"
#     fi
# }
# trap 'cleanup $LINENO "$BASH_COMMAND"' EXIT

# start an anvil instance in the background that has eigenlayer contracts deployed
# start_anvil_docker $parent_path/eigenlayer-deployed-anvil-state.json $parent_path/avs-and-eigenlayer-deployed-anvil-state.json

cd $parent_path/../../contracts
forge script script/IncredibleSquaringDeployer.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --legacy --broadcast -v
forge script script/RewardPoolDeployer.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --legacy --broadcast -v
# save the block-number in the genesis file which we also need to restart the anvil chain at the correct block
# otherwise the indexRegistry has a quorumUpdate at a high block number, and when we restart a clean anvil (without genesis.json) file
# it starts at block 0, and so calling getOperatorListAtBlockNumber reverts because it thinks there are no quorums registered at block 0
# EDIT: this doesn't actually work... since we can't both load a state and a genesis.json file... see https://github.com/foundry-rs/foundry/issues/6679
# will keep here in case this PR ever gets merged.
GENESIS_FILE=$parent_path/genesis.json
TMP_GENESIS_FILE=$parent_path/genesis.json.tmp
jq '.number = "'$(cast block-number)'"' $GENESIS_FILE >$TMP_GENESIS_FILE
mv $TMP_GENESIS_FILE $GENESIS_FILE

# we also do this here to make sure the operator has funds to register with the eigenlayer contracts
cast send 0x860B6912C2d0337ef05bbC89b0C2CB6CbAEAB4A5 --value 10ether --private-key 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6

token_address=$(jq -r '.addresses.erc20MockStrategy' "script/output/31337/credible_squaring_avs_deployment_output.json")
echo "Extracted token: $token_address"
registry_coordinator=$(jq -r '.addresses.registryCoordinator' "script/output/31337/credible_squaring_avs_deployment_output.json")
echo "Extracted registry_coordinator: $registry_coordinator"
operator_state_retriever=$(jq -r '.addresses.operatorStateRetriever' "script/output/31337/credible_squaring_avs_deployment_output.json")
echo "Extracted operator_state_retriever: $operator_state_retriever"

strategy_manager=$(jq -r '.addresses.strategyManager' "script/output/31337/eigenlayer_deployment_output.json")
echo "Extracted strategy_manager: $strategy_manager"

# whitelist newly deployed strategy
cast send "$strategy_manager" "addStrategiesToDepositWhitelist(address[],bool[])" "[$token_address]" "[true]" --private-key "$PRIVATE_KEY"

cd $parent_path/../../config-files
# File to update
OPERATOR_CONFIG_FILE="./operator.anvil.yaml"
DOCKER_OPERATOR_CONFIG_FILE="./operator-docker-compose.anvil.yaml"

# Update the values in the operator file
sed -i '' "s/avs_registry_coordinator_address: 0x[0-9a-fA-F]\+/avs_registry_coordinator_address: $registry_coordinator/" "$OPERATOR_CONFIG_FILE"
sed -i '' "s/operator_state_retriever_address: 0x[0-9a-fA-F]\+/operator_state_retriever_address: $operator_state_retriever/" "$OPERATOR_CONFIG_FILE"
sed -i '' "s/token_strategy_addr: 0x[0-9a-fA-F]\+/token_strategy_addr: $token_address/" "$OPERATOR_CONFIG_FILE"

# Update the values in the docker operator file
sed -i '' "s/avs_registry_coordinator_address: 0x[0-9a-fA-F]\+/avs_registry_coordinator_address: $registry_coordinator/" "$DOCKER_OPERATOR_CONFIG_FILE"
sed -i '' "s/operator_state_retriever_address: 0x[0-9a-fA-F]\+/operator_state_retriever_address: $operator_state_retriever/" "$DOCKER_OPERATOR_CONFIG_FILE"
sed -i '' "s/token_strategy_addr: 0x[0-9a-fA-F]\+/token_strategy_addr: $token_address/" "$DOCKER_OPERATOR_CONFIG_FILE"
