package main

type Configuration struct {
	Prod             string   `json:"prod"`
	Mnemonic         []string `json:"mnemonic"`
	SolanaPrivateKey string   `json:"Solana-PrivateKey"`
	BlockFrostAPIKey string   `json:"BlockFrost-API-Key"`
}

const Prod_Debug = "DEBUG"
