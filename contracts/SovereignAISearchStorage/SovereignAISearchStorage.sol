// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title SovereignAISearchStorage
 * @dev Smart contract for storing and retrieving user AI search data
 */
contract SovereignAISearchStorage {
    struct SearchResult {
        string query;           // Search query
        string resultHashCID;   // IPFS hash of AI search results
        uint256 timestamp;      // Search timestamp
        bool isEncrypted;       // Whether the result is encrypted
        string encryptionKey;   // Encryption key
        uint256 searchScore;    // Search relevance score
    }

    struct UserProfile {
        uint256 searchCount;    // Total number of searches by user
        bool isActive;          // User status
        uint256 lastSearchTime; // Last search timestamp
    }

    // Mapping from user address to user profile
    mapping(address => UserProfile) public userProfiles;
    
    // Mapping from user address to search history
    mapping(address => SearchResult[]) private userSearchHistory;
    
    // Event declarations
    event SearchResultStored(
        address indexed user,
        string query,
        uint256 timestamp
    );
    
    event UserProfileCreated(
        address indexed user,
        uint256 timestamp
    );

    // Modifier to check if user is active
    modifier onlyActiveUser() {
        require(userProfiles[msg.sender].isActive, "User not active");
        _;
    }

    /**
     * @dev Create a new user profile
     */
    function createUserProfile() external {
        require(!userProfiles[msg.sender].isActive, "User already exists");
        
        UserProfile storage newProfile = userProfiles[msg.sender];
        newProfile.isActive = true;
        newProfile.searchCount = 0;
        newProfile.lastSearchTime = block.timestamp;
        
        emit UserProfileCreated(msg.sender, block.timestamp);
    }

    /**
     * @dev Store a new search result
     * @param query Search query string
     * @param resultHashCID IPFS hash of results
     * @param isEncrypted Whether the result is encrypted
     * @param encryptionKey Optional encryption key
     * @param searchScore Relevance score of the search
     */
    function storeSearchResult(
        string memory query,
        string memory resultHashCID,
        bool isEncrypted,
        string memory encryptionKey,
        uint256 searchScore
    ) external onlyActiveUser {
        SearchResult memory newResult = SearchResult({
            query: query,
            resultHashCID: resultHashCID,
            timestamp: block.timestamp,
            isEncrypted: isEncrypted,
            encryptionKey: encryptionKey,
            searchScore: searchScore
        });
        
        userSearchHistory[msg.sender].push(newResult);
        
        UserProfile storage profile = userProfiles[msg.sender];
        profile.searchCount++;
        profile.lastSearchTime = block.timestamp;
        
        emit SearchResultStored(msg.sender, query, block.timestamp);
    }

    /**
     * @dev Get paginated search history for user
     * @param pageSize Number of results per page
     * @param pageNumber Page number to retrieve
     * @return SearchResult[] Array of search results
     */
    function getSearchHistory(uint256 pageSize, uint256 pageNumber) 
        external 
        view 
        onlyActiveUser 
        returns (SearchResult[] memory) 
    {
        require(pageSize > 0, "Page size must be greater than 0");
        
        SearchResult[] storage history = userSearchHistory[msg.sender];
        uint256 startIndex = pageSize * pageNumber;
        
        require(startIndex < history.length, "Page number out of range");
        
        uint256 endIndex = startIndex + pageSize;
        if (endIndex > history.length) {
            endIndex = history.length;
        }
        
        uint256 resultLength = endIndex - startIndex;
        SearchResult[] memory results = new SearchResult[](resultLength);
        
        for (uint256 i = 0; i < resultLength; i++) {
            results[i] = history[startIndex + i];
        }
        
        return results;
    }

    /**
     * @dev Get most recent search results
     * @param count Number of recent results to retrieve
     */
    function getRecentSearches(uint256 count) 
        external 
        view 
        onlyActiveUser 
        returns (SearchResult[] memory) 
    {
        require(count > 0, "Count must be greater than 0");
        
        SearchResult[] storage history = userSearchHistory[msg.sender];
        uint256 resultLength = count;
        
        if (count > history.length) {
            resultLength = history.length;
        }
        
        SearchResult[] memory results = new SearchResult[](resultLength);
        for (uint256 i = 0; i < resultLength; i++) {
            results[i] = history[history.length - resultLength + i];
        }
        
        return results;
    }

    /**
     * @dev Get user's search statistics
     */
    function getUserStats() 
        external 
        view 
        onlyActiveUser 
        returns (
            uint256 totalSearches,
            uint256 lastSearchTime
        ) 
    {
        UserProfile storage profile = userProfiles[msg.sender];
        return (
            profile.searchCount,
            profile.lastSearchTime
        );
    }

    /**
     * @dev Search history by keyword
     * @param keyword Search keyword to match
     */
    function searchHistoryByKeyword(string memory keyword) 
        external 
        view 
        onlyActiveUser 
        returns (SearchResult[] memory) 
    {
        require(bytes(keyword).length > 0, "Keyword cannot be empty");
        
        SearchResult[] storage history = userSearchHistory[msg.sender];
        uint256 matchCount = 0;
        
        // First count matching results
        for (uint256 i = 0; i < history.length; i++) {
            if (contains(history[i].query, keyword)) {
                matchCount++;
            }
        }
        
        // Create results array
        SearchResult[] memory results = new SearchResult[](matchCount);
        uint256 resultIndex = 0;
        
        // Fill results array
        for (uint256 i = 0; i < history.length; i++) {
            if (contains(history[i].query, keyword)) {
                results[resultIndex] = history[i];
                resultIndex++;
            }
        }
        
        return results;
    }

    /**
     * @dev Check if a string contains a substring
     * @param source Source string to search in
     * @param search Substring to search for
     * @return bool True if substring is found
     */
    function contains(string memory source, string memory search) 
        internal 
        pure 
        returns (bool) 
    {
        bytes memory sourceBytes = bytes(source);
        bytes memory searchBytes = bytes(search);

        if (searchBytes.length == 0 || sourceBytes.length < searchBytes.length) {
            return false;
        }

        for (uint i = 0; i <= sourceBytes.length - searchBytes.length; i++) {
            bool found = true;
            for (uint j = 0; j < searchBytes.length; j++) {
                if (sourceBytes[i + j] != searchBytes[j]) {
                    found = false;
                    break;
                }
            }
            if (found) {
                return true;
            }
        }
        return false;
    }
}