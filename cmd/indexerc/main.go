package main

import (
	"context"
	"fmt"
	"os"

	"github.com/GregoireBailly/indexerc/internal/eth"
	"github.com/GregoireBailly/indexerc/internal/analyzer"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("🚀 Starting indexERC — connecting to Ethereum mainnet…")

    url := os.Getenv("ETH_RPC_URL")
    key := os.Getenv("ETH_RPC_API_KEY")
    if url == "" || key == "" {
        return fmt.Errorf("missing connexion info in environment")
	}

    rpcURL := fmt.Sprintf("%s/%s", url, key)
	ctx := context.Background()

	client, err := eth.New(ctx, rpcURL)
	if err != nil {
		return fmt.Errorf("Failed to connect to Ethereum: %v\n", err)
	}
	defer client.Close()

	counter := analyzer.NewCounter(client)

	count, err := counter.Count(ctx, 10, eth.ERC20TransferQuery())
	if err != nil {
		return fmt.Errorf("Failed to count: %v\n", err)
	}

	fmt.Printf("✅ Connected!\n")
	fmt.Printf("Count of ERC20 tranfer is: %d\n", count)
	return nil
}
