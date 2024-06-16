package dockerfileParser

import (
  "bufio"
	"fmt"
	"os"
	"strings"
	"encoding/json"
)

func ParseDockerfile(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	var instructions []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			instructions = append(instructions, line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return instructions
}

func ExecuteInstruction(instruction string) {
	parts := strings.Fields(instruction)

	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "FROM":
		fmt.Println("Handling FROM instruction")
		if len(parts) != 2 {
			fmt.Println("Invalid FROM instruction")
			return
		}
		baseImage := parts[1]
		installNode(baseImage)
	case "WORKDIR":
		if len(parts) != 2 {
			fmt.Println("Invalid WORKDIR instruction")
			return
		}
		dir := parts[1]
		changeWorkDir(dir)
	case "COPY":
		fmt.Println("Handling COPY instruction")
		if len(parts) != 3 {
			fmt.Println("Invalid COPY instruction")
			return
		}
		src := parts[1]
		dst := parts[2]
		copyFile(src, dst)
	case "RUN":
		fmt.Println("Handling RUN instruction")
		if len(parts) < 2 {
			fmt.Println("Invalid RUN instruction")
			return
		}
    runCommand(parts[1:])
	case "CMD":
		fmt.Println("Handling CMD instruction")
		cmdString := strings.Join(parts[1:], " ")
		var cmdArray []string
		if err := json.Unmarshal([]byte(cmdString), &cmdArray); err != nil {
			fmt.Printf("Error parsing CMD instruction: %v\n", err)
			return
		}
		runApplication(cmdArray)
	default:
		fmt.Printf("Unknown instruction: %s\n", parts[0])
	}
}
