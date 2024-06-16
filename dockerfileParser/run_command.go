package dockerfileParser

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
)

func runCommand(commandParts []string) {
  fmt.Printf("Running command: %s\n", strings.Join(commandParts, " "))

	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running command %s: %v\n", strings.Join(commandParts, " "), err)
	}
}
