package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/BlockCraftsman/SovereignSearch-oasis-Protocol/connection"
	"github.com/BlockCraftsman/SovereignSearch-oasis-Protocol/contracts/SovereignAISearchStorage"
)

var createUserProfileCmd = &cobra.Command{
	Use:   "createUserProfile [flags] contract_address",
	Short: "Activate a user profile in the SovereignAISearchStorage contract",
	Args:  cobra.ExactArgs(1),
	Run:   CreateUserProfile,
}

func init() {
	RootCmd.AddCommand(createUserProfileCmd)
}

func CreateUserProfile(cmd *cobra.Command, args []string) {
	// Set up a context for calls with a timeout of 1 minute.
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancelCtx()

	contractAddr, err := ParseAddress(args[0])
	if err != nil {
		ExitWithError("Unable to parse contract address", err)
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

	fmt.Fprintf(os.Stderr, "Activating user profile in SovereignAISearchStorage...\n")

	// Create the user profile.
	createUserProfileTx, err := sas.CreateUserProfile(auth)
	if err != nil {
		ExitWithError("Failed to create user profile tx", err)
	}

	_, err = bind.WaitMined(ctx, conn.Sapphire, createUserProfileTx)
	if err != nil {
		ExitWithError("Failed to activate user profile", err)
	}

	fmt.Fprintf(os.Stderr, "User profile activated successfully.\n")
}
