package Cardano

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type GenesisVerificationKey struct {
	AddressType     string `json:"type"`
	Description     string `json:"description"`
	VerificationKey string `json:"cborHex"`
}

type GenesisSigningKey struct {
	AddressType string `json:"type"`
	Description string `json:"description"`
	SigningKey  string `json:"cborHex"`
}

func GenesisCommand(args ...string) {
	var command []string

	command = append(command, "genesis")
	command = append(command, args...)

	RunCli(command...)
}

func CreateGenesisKeyPair() (GenesisVerificationKey, GenesisSigningKey) {
	GenesisCommand("key-gen-genesis", "--verification-key-file", "verification.json", "--signing-key-file", "signing.json")

	verificationFile, _ := ioutil.ReadFile("verification.json")
	verificationKey := GenesisVerificationKey{}
	json.Unmarshal([]byte(verificationFile), &verificationKey)

	signingFile, _ := ioutil.ReadFile("signing.json")
	signingKey := GenesisSigningKey{}
	json.Unmarshal([]byte(signingFile), &signingKey)

	os.Remove("verification.json")
	os.Remove("signing.json")

	return verificationKey, signingKey
}
