package model

// Configs is a struct that holds the blockchain configuration for onchain resources
// A key has a config.
// Config has a chain id, a contract address, a contract ABI

type Config struct {
	ID string `json:"id"`
	ChainID string `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
	ContractABI string `json:"contract_abi"`
	Confirmed bool `json:"confirmed"`
}

