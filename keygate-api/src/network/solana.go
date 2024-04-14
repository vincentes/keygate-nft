package network

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/blocto/solana-go-sdk/client"
)

var (
	solanaClient *client.Client
	once         sync.Once
)

func GetSolanaClient() *client.Client {
	once.Do(func() {
		solanaClient = client.NewClient("https://api.devnet.solana.com")

		resp, err := solanaClient.GetVersion(context.Background())
		if err != nil {
			log.Fatalf("Failed to get version: %v", err)
		}
		
		fmt.Printf("Version: %v\n", resp)
	})
	return solanaClient
}