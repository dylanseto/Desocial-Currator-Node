package Cardano

import (
	"fmt"
	"os/exec"
	"runtime"
)

const i = 100

func RunCli(args string) {
	switch runtime.GOOS {
	case "windows":
		mCmd := "./lib/cardano-cli.exe"
		cmd, err := exec.Command(mCmd, args).Output()

		if err != nil {
			fmt.Println("error: ", err)
		}

		fmt.Println(string(cmd))
		break
	}
}
