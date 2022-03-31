package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/martinlindhe/notify"
)

func main() {
	application := os.Args[1]
	args := []string{}
	if len(os.Args) > 2 {
		for i := 1; i < len(os.Args); i++ {
			args = append(args, os.Args[i])
		}
	}

	cmd := exec.Command(application, args...)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Print(err)
	}
	notify.Notify("CLI Notify", "Command finished", application, "cmd.png")
}
