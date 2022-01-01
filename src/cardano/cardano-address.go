package Cardano

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ShelleyVerificationKey struct {
	AddressType     string `json:"type"`
	Description     string `json:"description"`
	VerificationKey string `json:"cborHex"`
}

type ShelleySigningKey struct {
	AddressType string `json:"type"`
	Description string `json:"description"`
	SigningKey  string `json:"cborHex"`
}

func AddressCommand(args ...string) {
	var command []string

	command = append(command, "address")
	command = append(command, args...)

	RunCli(command...)
}

func AddressHelpCommand() {
	AddressCommand("--help")
}

func GenerateShelleyKey() (ShelleyVerificationKey, ShelleySigningKey) {
	AddressCommand("key-gen", "--normal-key", "--verification-key-file", "verification.json", "--signing-key-file", "signing.json")

	verificationFile, _ := ioutil.ReadFile("verification.json")
	verificationKey := ShelleyVerificationKey{}
	json.Unmarshal([]byte(verificationFile), &verificationKey)

	signingFile, _ := ioutil.ReadFile("signing.json")
	signingKey := ShelleySigningKey{}
	json.Unmarshal([]byte(signingFile), &signingKey)

	os.Remove("verification.json")
	os.Remove("signing.json")

	return verificationKey, signingKey
}
