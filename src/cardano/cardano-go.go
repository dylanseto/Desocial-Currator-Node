package Cardano

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func RunCli(args ...string) {
	switch runtime.GOOS {
	case "windows":
		mCmd := "./lib/cardano-cli"
		cmd := exec.Command(mCmd, args...)

		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()

		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return
		}
		fmt.Println("Result: " + out.String())

		break
	default:
		fmt.Println("This Operating System is not supported!")
	}
}

func ShowHelp() {
	RunCli("help")
}

func ShowVersion() {
	RunCli("version")
}
