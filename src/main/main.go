package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var config Configuration

func loadConf() error {
	file, _ := ioutil.ReadFile("conf/conf.json")
	config = Configuration{}
	err := json.Unmarshal([]byte(file), &config)
	return err
}

func main() {
	fmt.Print("[Event] Loading Configuration File...")
	err := loadConf()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done!")
	}

	fmt.Print("[Event] Connecting To Solana...")
	account, err := solana.WalletFromPrivateKeyBase58(config.SolanaPrivateKey)

	if err != nil {
		// Private Key is invalid
		fmt.Println(err)
		return
	}

	fmt.Println("Done!")

	if config.Prod == Prod_Debug {

		client := rpc.New(rpc.TestNet_RPC)

		balance, err := client.GetBalance(
			context.TODO(),
			account.PublicKey(),
			rpc.CommitmentFinalized,
		)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("balance:", balance.Value)
	}

	fmt.Print("[Event] Connecting To Cardano (Blockfrost)...")

	cardano.runCli()
}
