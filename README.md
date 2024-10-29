# SovereignAISearch-oasis-Protocol

## Project Description
SovereignAISearch-oasis-Protocol is a decentralized AI search protocol designed to provide users with a secure, privacy-preserving, and personalized search experience. The protocol leverages smart contracts and blockchain technology to ensure that user search data is not collected and exploited by centralized entities, while delivering efficient and AI-driven search results.

SovereignAISearch-oasis-Protocol represents a paradigm shift in search technology, prioritizing user data sovereignty and privacy through smart contracts and blockchain technology. This next-generation search engine empowers users with control over their information and offers an efficient, secure, and AI-driven search solution.

## Key Features
- **Decentralized AI Search**: Utilizes blockchain technology to provide a trustless AI search service, ensuring the security and privacy of user data.
- **Smart Contract Support**: Manages search requests and results through smart contracts, ensuring transparent and tamper-proof search records.
- **User CIDS Management**: Users can add and manage their own CIDS (Content Identifier Strings) and use these identifiers in their searches.
- **Multi-Author Support**: Supports multiple authors to set and view their own messages, ensuring the diversity and credibility of information.
- **Advanced AI Capabilities**: Leverages AI to analyze user behavior and preferences, offering highly personalized search results.
- **Decentralized Storage**: Ensures user data sovereignty by storing search data across a decentralized network using the Storcha platform.


## Smart Contract Architecture

### Overview
The `SovereignAISearchStorage` smart contract is designed to manage user profiles and search history securely and efficiently. The contract ensures that user data is stored and retrieved in a decentralized manner, maintaining user privacy and data sovereignty.

### Architecture Diagram
```mermaid
graph TB
    subgraph SovereignAISearchStorage[SovereignAISearchStorage Contract]
        
        subgraph DataStructures[Data Structures]
            SearchResult[SearchResult Struct<br>- query<br>- resultHashCID<br>- timestamp<br>- isEncrypted<br>- encryptionKey<br>- searchScore]
            UserProfile[UserProfile Struct<br>- searchCount<br>- isActive<br>- lastSearchTime]
        end

        subgraph Storage[Storage]
            Mappings[Storage Mappings]
            Mappings --> UserProfiles[userProfiles<br>address => UserProfile]
            Mappings --> SearchHistory[userSearchHistory<br>address => SearchResult[]]
        end

        subgraph Functions[Public Functions]
            CreateProfile[createUserProfile]
            StoreSearch[storeSearchResult]
            GetHistory[getSearchHistory]
            GetRecent[getRecentSearches]
            GetStats[getUserStats]
            SearchKeyword[searchHistoryByKeyword]
        end

        subgraph Internal[Internal Functions]
            Contains[contains<br>String Search Helper]
        end

        subgraph Events[Events]
            StoredEvent[SearchResultStored]
            ProfileEvent[UserProfileCreated]
        end

    end

    User((User)) --> CreateProfile
    User --> StoreSearch
    User --> GetHistory
    User --> GetRecent
    User --> GetStats
    User --> SearchKeyword

    CreateProfile --> UserProfiles
    StoreSearch --> SearchHistory
    GetHistory --> SearchHistory
    GetRecent --> SearchHistory
    GetStats --> UserProfiles
    SearchKeyword --> SearchHistory
    SearchKeyword --> Contains
```

### Detailed Architecture

#### UserProfile Mapping
- **UserProfile**: Stores user-specific information such as search count, user status, and last search timestamp.
- **Mapping**: Maps user addresses to their respective `UserProfile` structs.

#### SearchResult Mapping
- **SearchResult**: Stores individual search results, including query, result hash CID, timestamp, encryption status, encryption key, and search score.
- **Mapping**: Maps user addresses to arrays of `SearchResult` structs, representing the user's search history.

#### Events
- **SearchResultStored**: Emitted when a new search result is stored, capturing the user address, query, and timestamp.
- **UserProfileCreated**: Emitted when a new user profile is created, capturing the user address and timestamp.

#### Modifiers
- **onlyActiveUser**: Ensures that only active users can perform certain actions, such as storing search results or retrieving search history.

#### Functions
- **createUserProfile**: Allows users to create a new profile, initializing their search count, status, and last search timestamp.
- **storeSearchResult**: Enables users to store new search results, updating their search history and profile statistics.
- **getSearchHistory**: Retrieves paginated search history for the user, based on page size and page number.
- **getRecentSearches**: Retrieves the most recent search results for the user, based on the specified count.
- **getUserStats**: Retrieves the user's search statistics, including total searches and last search timestamp.
- **searchHistoryByKeyword**: Searches the user's search history for results containing a specific keyword.
- **contains**: A helper function to check if a string contains a substring, used for keyword search functionality.

### Workflow Description

1. **User Profile Creation**:
   - The user invokes the `createUserProfile` function to initialize their profile.
   - The contract creates a new `UserProfile` struct for the user, setting their status to active and initializing their search count and last search timestamp.

2. **Storing Search Results**:
   - The user invokes the `storeSearchResult` function to store a new search result.
   - The contract appends the new `SearchResult` to the user's search history and updates their profile statistics.

3. **Retrieving Search History**:
   - The user invokes the `getSearchHistory` or `getRecentSearches` function to retrieve their search history.
   - The contract retrieves the requested search results from the user's search history and returns them.

4. **Searching by Keyword**:
   - The user invokes the `searchHistoryByKeyword` function to search their search history for results containing a specific keyword.
   - The contract iterates through the user's search history, checking each query for the keyword, and returns matching results.

5. **Retrieving User Statistics**:
   - The user invokes the `getUserStats` function to retrieve their search statistics.
   - The contract retrieves the user's profile and returns their total search count and last search timestamp.

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