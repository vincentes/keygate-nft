package network

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/types"
)

var (
	solanaClient *client.Client
	mode         = "dev"
	once         sync.Once
	systemAccount   types.Account
)

// map mode to network
var networkMap = map[string]string{
	"dev":  "https://api.devnet.solana.com",
	"test": "https://api.testnet.solana.com",
	"main": "https://api.mainnet-beta.solana.com",
}

func GetSolanaClient() *client.Client {
	once.Do(func() {
		solanaClient = client.NewClient(networkMap[mode])
		resp, err := solanaClient.GetVersion(context.Background())
		if err != nil {
			log.Fatalf("Failed to get version: %v", err)
		}
		log.Printf("Connected to Solana cluster: %s", resp.SolanaCore)

		privateKey := os.Getenv("SYSTEM_ACCOUNT_PK")
		if privateKey == "" {
			log.Fatal("SYSTEM_ACCOUNT_PK environment variable is not set")
		}

		systemAccount, _ = types.AccountFromBase58(privateKey)
		log.Printf("System wallet loaded: %s", systemAccount.PublicKey.String())
	})
	return solanaClient
}

func GetSystemAccount() types.Account {
	return systemAccount
}