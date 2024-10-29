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

var storeCIDCmd = &cobra.Command{
	Use:   "storeCID [flags] contract_address query resultHashCID isEncrypted encryptionKey searchScore",
	Short: "Store the CID information inside the SovereignAISearchStorage contract",
	Args:  cobra.ExactArgs(6),
	Run:   StoreCID,
}

func init() {
	RootCmd.AddCommand(storeCIDCmd)
}

func StoreCID(cmd *cobra.Command, args []string) {
	// Set up a context for calls with a timeout of 1 minute.
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancelCtx()

	contractAddr, err := ParseAddress(args[0])
	if err != nil {
		ExitWithError("Unable to parse contract address", err)
	}

	query := args[1]
	resultHashCID := args[2]
	isEncrypted := args[3] == "true"
	encryptionKey := args[4]
	searchScore := new(big.Int)
	_, ok := searchScore.SetString(args[5], 10)
	if !ok {
		ExitWithError("Unable to parse search score", fmt.Errorf("invalid search score: %s", args[5]))
	}

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

	// Prepare the transaction.
	auth, err := conn.PrepareNextTx(ctx)
	if err != nil {
		ExitWithError("Failed to prepare next tx", err)
	}

	fmt.Fprintf(os.Stderr, "Storing CID information in SovereignAISearchStorage...\n")

	// Store the search result.
	storeCIDTx, err := sas.StoreSearchResult(auth, query, resultHashCID, isEncrypted, encryptionKey, searchScore)
	if err != nil {
		ExitWithError("Failed to create store tx", err)
	}

	_, err = bind.WaitMined(ctx, conn.Sapphire, storeCIDTx)
	if err != nil {
		ExitWithError("Failed to store CID information", err)
	}

	fmt.Fprintf(os.Stderr, "CID information stored successfully.\n")
}
