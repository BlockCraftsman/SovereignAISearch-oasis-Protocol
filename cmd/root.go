package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

const (
	// NetworkFlag is the name of the flag used to specify the network.
	NetworkFlag = "network"
)

var (
	// RootCmd is the root command for the SovereignAISearchStorage application.
	RootCmd = &cobra.Command{
		Use:   "SovereignAISearchStorage",
		Short: "SovereignAISearch empowers users with complete control over AI search results, ensuring decentralized storage and privacy.",
	}
)

func init() {
	// Add a persistent flag to specify the network.
	RootCmd.PersistentFlags().String(NetworkFlag, "sapphire-localnet", "Name of the network to connect to")
}

// ExitWithError terminates the program after writing the error to stderr.
func ExitWithError(msg string, err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s: %v\n", msg, err)
	os.Exit(1)
}

// GetNetworkAddress returns the dial address of the network specified via the
// network flag. If the network is unknown, it exits with an error.
func GetNetworkAddress() string {
	networks := map[string]string{
		"sapphire":          "https://sapphire.oasis.io",
		"sapphire-testnet":  "https://testnet.sapphire.oasis.io",
		"sapphire-localnet": "http://localhost:8545",
	}

	net, err := RootCmd.PersistentFlags().GetString(NetworkFlag)
	if err != nil {
		ExitWithError("GetNetworkAddress", fmt.Errorf("Please specify the network to connect to using --%s.", NetworkFlag))
	}
	net = strings.ToLower(net)

	addr, found := networks[net]
	if !found {
		validNets := make([]string, 0, len(networks))
		for n := range networks {
			validNets = append(validNets, n)
		}

		ExitWithError("GetNetworkAddress", fmt.Errorf("Unknown network specified, please use one of the following: %s.", strings.Join(validNets, ", ")))
	}

	return addr
}

// ParseAddress converts the hex representation of an Ethereum address into
// common.Address. It returns an error if the address is malformed.
func ParseAddress(addrHex string) (common.Address, error) {
	if strings.HasPrefix(addrHex, "0x") {
		addrHex = strings.TrimPrefix(addrHex, "0x")
	}

	if len(addrHex) != 40 {
		return common.Address{}, fmt.Errorf("address is malformed")
	}

	return common.HexToAddress(addrHex), nil
}

// Execute runs the root command and handles any errors.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		ExitWithError("SovereignAISearchStorage", err)
	}
}
