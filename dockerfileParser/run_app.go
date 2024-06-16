package dockerfileParser

import (
  "fmt"
  "os"
  "os/exec"
)

func runApplication(commandParts []string) {
	if len(commandParts) == 0 {
		fmt.Println("No command provided to run the application")
		return
	}

	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting application: %v\n", err)
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("Application exited with error: %v\n", err)
		return
	}

	fmt.Println("Application exited successfully.")
}
