package cmd

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/BlockCraftsman/SovereignAISearch-oasis-Protocol/connection"
	"github.com/BlockCraftsman/SovereignAISearch-oasis-Protocol/contracts/SovereignAISearchStorage"
)

var getAIQueryCmd = &cobra.Command{
	Use:   "getCID [flags] contract_address cid",
	Short: "Get the message stored inside the SovereignAISearchStorage contract for a specific CID",
	Args:  cobra.ExactArgs(2),
	Run:   GetAIQueryCIDContent,
}

func init() {
	RootCmd.AddCommand(getAIQueryCmd)
}

func GetAIQueryCIDContent(cmd *cobra.Command, args []string) {
	// Set up a context for calls with a timeout of 1 minute.
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancelCtx()

	contractAddr, err := ParseAddress(args[0])
	if err != nil {
		ExitWithError("Unable to parse contract address", err)
	}

	cid := args[1]

	// Connect to the network.
	conn, err := connection.NewConnection(ctx, GetNetworkAddress())
	if err != nil {
		ExitWithError("Unable to connect", err)
	}

	// Get an instance of the SovereignAISearchStorage contract.
	sas, err := SovereignAISearchStorage.NewSovereignAISearchStorage(contractAddr, conn.Sapphire)
	if err != nil {
		ExitWithError("Unable to get instance of contract", err)
	}

	// Retrieve search history for the given CID.
	fmt.Fprintf(os.Stderr, "Retrieving search history for CID %s...\n", cid)

	// Define the page size and page number for the search history.
	pageSize := big.NewInt(10)  // Example page size
	pageNumber := big.NewInt(0) // Example page number

	searchResults, err := sas.GetSearchHistory(&bind.CallOpts{From: conn.Address}, pageSize, pageNumber)
	if err != nil {
		ExitWithError("Failed to retrieve search history", err)
	}

	// Filter the search results for the specific CID.
	var foundResult *SovereignAISearchStorage.SovereignAISearchStorageSearchResult
	for _, result := range searchResults {
		if result.ResultHashCID == cid {
			foundResult = &result
			break
		}
	}

	if foundResult == nil {
		fmt.Fprintf(os.Stderr, "No search result found for CID %s\n", cid)
		return
	}

	// Output retrieved message to stdout.
	fmt.Printf("Query: %s\nResultHashCID: %s\nTimestamp: %s\nIsEncrypted: %t\nEncryptionKey: %s\nSearchScore: %s\n",
		foundResult.Query, foundResult.ResultHashCID, foundResult.Timestamp.String(), foundResult.IsEncrypted, foundResult.EncryptionKey, foundResult.SearchScore.String())
}
