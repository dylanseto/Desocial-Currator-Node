package Cardano

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

type AddressInfo struct {
	Address       string `json:"address"`
	Era           string `json:"era"`
	Encoding      string `json:"encoding"`
	Type          string `json:"type"`
	AddressBase16 string `json:"base16"`
}

func AddressCommand(args ...string) string {
	var command []string

	command = append(command, "address")
	command = append(command, args...)

	res := RunCli(command...)
	return res
}

func AddressHelpCommand() {
	AddressCommand("--help")
}

func GenerateShelleyPaymentKey() (ShelleyVerificationKey, ShelleySigningKey) {
	AddressCommand("key-gen", "--normal-key", "--verification-key-file", "verification.json", "--signing-key-file", "signing.json")

	verificationFile, _ := ioutil.ReadFile("verification.json")
	verificationKey := ShelleyVerificationKey{}
	json.Unmarshal([]byte(verificationFile), &verificationKey)

	signingFile, _ := ioutil.ReadFile("signing.json")
	signingKey := ShelleySigningKey{}
	json.Unmarshal([]byte(signingFile), &signingKey)

	os.Remove("verification.json")
	os.Remove("signing.json")
	verificationKeyPrefix := "addr_vk"
	fmt.Println(verificationKeyPrefix + verificationKey.VerificationKey)

	return verificationKey, signingKey
}

func BuildPaymentAddress(verificationKey ShelleyVerificationKey, mainnet bool) string {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}

	txt, _ := json.Marshal(verificationKey)
	tmpFile.Write(txt)

	if mainnet {
		AddressCommand("build", "--payment-verification-key-file", tmpFile.Name(), "--out-file", "payment.addr", "--mainnet")
	} else {
		AddressCommand("build", "--payment-verification-key-file", tmpFile.Name(), "--out-file", "payment.addr", "--testnet-magic", "1097911063")
	}

	os.Remove(tmpFile.Name())

	paymentAddress, _ := ioutil.ReadFile("payment.addr")

	return string(paymentAddress)
}

func GetPaymentAddressInfo(address string) AddressInfo {
	res := AddressCommand("info", "--address", address)
	addressInfo := AddressInfo{}

	json.Unmarshal([]byte(res), &addressInfo)

	return addressInfo
}

func GetPaymentAddressKeyHash(vk string) string {
	res := AddressCommand("key-hash", "--payment-verification-key", vk)

	return res
}
