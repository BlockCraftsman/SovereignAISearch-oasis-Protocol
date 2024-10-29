# SovereignAISearch-oasis-Protocol

## Project Description
SovereignAISearch-oasis-Protocol is a decentralized AI search protocol designed to provide users with a secure, privacy-preserving, and personalized search experience. The protocol leverages smart contracts and blockchain technology to ensure that user search data is not collected and exploited by centralized entities, while delivering efficient and AI-driven search results.

## Key Features
- **Decentralized AI Search**: Utilizes blockchain technology to provide a trustless AI search service, ensuring the security and privacy of user data.
- **Smart Contract Support**: Manages search requests and results through smart contracts, ensuring transparent and tamper-proof search records.
- **User CIDS Management**: Users can add and manage their own CIDS (Content Identifier Strings) and use these identifiers in their searches.
- **Multi-Author Support**: Supports multiple authors to set and view their own messages, ensuring the diversity and credibility of information.
- **Advanced AI Capabilities**: Leverages AI to analyze user behavior and preferences, offering highly personalized search results.
- **Decentralized Storage**: Ensures user data sovereignty by storing search data across a decentralized network using the Storcha platform.

## Prerequisites
To build and test the SovereignAISearch-oasis-Protocol, you need the following tools:
- `solc` (Solidity compiler)
- `abigen` (Go Ethereum ABI generator)
- Go (version 1.16 or later)

If you're on Ubuntu, you can install `solc` and `abigen` with:
```shell
make install-deps
```

## Building
To build the project, simply run:
```shell
make
```

## Testing
To run the end-to-end test, you should first start the Sapphire Localnet in Docker:
```shell
make run-localnet
```

After it has finished starting up, you can run the end-to-end test:
```shell
make test
```

This will deploy the contract, activate a user profile, store a CID (Content Identifier) in it, retrieve it, and check if the retrieved CID matches the stored one.

## Running
Before deploying the contract or interacting with it, you should store your account's hex-encoded private key in an environment variable (the `0x` prefix is optional):
```shell
export PRIVATE_KEY=...
```

### Deploying the Contract
You can deploy the contract on different networks by invoking:
```shell
./SovereignSearch-oasis-Protocol deploy --network sapphire-localnet # Sapphire Localnet
./SovereignSearch-oasis-Protocol deploy --network sapphire-testnet  # Sapphire Testnet
./SovereignSearch-oasis-Protocol deploy --network sapphire          # Sapphire Mainnet
```

The deployed contract's address is printed to the standard output if the deployment is successful. You can store it in an environment variable, as you will need it to interact with the contract.

### Activating User Profile
To activate a user profile in the deployed contract:
```shell
./SovereignSearch-oasis-Protocol createUserProfile --network sapphire-localnet ${CONTRACT_ADDR}
```

### Storing CID Information
To store CID information inside the deployed contract:
```shell
./SovereignSearch-oasis-Protocol storeCID --network sapphire-localnet ${CONTRACT_ADDR} "do you know oasisprotocol web3 project?" "bafybeiekjfdzkn3bk4uzs3zlpuzdrmqyrht5joehnqvr6r7shlde5fbe7e" "false" "" "100"
```

### Retrieving CID Information
To retrieve the CID information from the deployed contract:
```shell
./SovereignSearch-oasis-Protocol getCID --network sapphire-localnet ${CONTRACT_ADDR} "bafybeiekjfdzkn3bk4uzs3zlpuzdrmqyrht5joehnqvr6r7shlde5fbe7e"
```

## Debugging
For debugging purposes, you can also run the localnet in debug mode with:
```shell
make run-localnet-debug
```

## Contributing
We welcome contributions to SovereignAISearch-oasis-Protocol! Please read our [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute to the project.

## License
SovereignAISearch-oasis-Protocol is licensed under the [Apache 2.0 License](LICENSE).