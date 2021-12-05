package main

type Configuration struct {
	Prod             string   `json:"prod"`
	Mnemonic         []string `json:"mnemonic"`
	SolanaPrivateKey string   `json:"Solana-PrivateKey"`
}

const Prod_Debug = "DEBUG"
