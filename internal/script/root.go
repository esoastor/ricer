package script

import (
	"fmt"
	"os"
	"os/exec"
	"ricer/internal/config"
)

func Run() {
	config := config.Get()
	commandParts := config.AfterCommand

	commandPartsLen := len(commandParts)
	if commandPartsLen == 0 {
		return
	}
	commandArgs := []string{}
	if commandPartsLen > 1 {
		commandArgs = commandParts[1:]
	}
	command := commandParts[0]

	cmd := exec.Command(command, commandArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Werror while running script: %s", err.Error())
		os.Exit(1)
	}
}
