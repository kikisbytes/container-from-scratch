package dockerfileParser

import (
  "fmt"
  "os"
)

func changeWorkDir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Printf("Error changing directory to %s: %v\n", dir, err)
		return
	}
	fmt.Printf("Successfully changed directory to %s\n", dir)
}
