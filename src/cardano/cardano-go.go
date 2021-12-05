package cardano

import (
	"os/exec"
	"runtime"
)

func runCli() {
	switch runtime.GOOS {
	case "windows":
		exec.Command("", "lib/cardano-cli.exe")
		break
	}
}
