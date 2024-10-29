#!/bin/sh

set -eu

# Private key of the first test account.
export PRIVATE_KEY="ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

# Network to use.
NETWORK="sapphire-localnet"

# Query to store in the contract.
QUERY="example query"

# Result Hash CID to store in the contract.
RESULT_HASH_CID="example_cid"

# Whether the result is encrypted.
IS_ENCRYPTED="false"

# Encryption key (if the result is encrypted).
ENCRYPTION_KEY=""

# Search score.
SEARCH_SCORE="100"

# Deploy the contract, its address is returned on stdout.
ADDR=$(./SovereignSearch-oasis-Protocol deploy --network ${NETWORK})
echo "Contract address is: ${ADDR}"

# Activate the user profile in the deployed contract.
./SovereignSearch-oasis-Protocol createUserProfile --network ${NETWORK} "${ADDR}"

# Store the CID information inside the deployed contract.
./SovereignSearch-oasis-Protocol storeCID --network ${NETWORK} "${ADDR}" "${QUERY}" "${RESULT_HASH_CID}" "${IS_ENCRYPTED}" "${ENCRYPTION_KEY}" "${SEARCH_SCORE}"

# Retrieve the CID information from the deployed contract.
CID_GOT=$(./SovereignSearch-oasis-Protocol getCID --network ${NETWORK} "${ADDR}" "${RESULT_HASH_CID}" | grep -oP 'ResultHashCID: \K.*')

# Check if the retrieved CID information matches the stored CID information.
if [ "x${RESULT_HASH_CID}" = "x${CID_GOT}" ]; then
	echo "Test passed!"
	echo "Stored \"${RESULT_HASH_CID}\", got \"${CID_GOT}\"."
	exit 0
else
	echo "Test failed!"
	echo "Expected \"${RESULT_HASH_CID}\", got \"${CID_GOT}\"."
	exit 1
fi