package dockerfileParser

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func installNode(baseImage string) {
  configResolvConf()

  nodeVersion := strings.Split(baseImage, ":")[1]

  fmt.Println("Installing Node.js version " + nodeVersion + "...")

  installNodeJSUsingApt(nodeVersion)
}

func configResolvConf() {
	content := `nameserver 127.0.0.53
options edns0 trust-ad
search home
`

	fmt.Println("Setting up resolv.conf to install Node.js...")

	err := os.WriteFile("/etc/resolv.conf", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing /etc/resolv.conf:", err)
	}
}

func installNodeJSUsingApt(version string) {
	cmd := exec.Command("")

  if checkIfNodeJSIsInstalled(version) {
    return
  }

  fmt.Printf("Installing Node.js version %s using apt...\n", version)

	if strings.HasPrefix(version, "v") {
		version = version[1:]
	}

	setupScriptURL := fmt.Sprintf("https://deb.nodesource.com/setup_%s.x", version)
	cmd = exec.Command("sh", "-c", fmt.Sprintf("curl -fsSL %s | sudo -E bash -", setupScriptURL))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error adding NodeSource repository: %v\n", err)
		return
	}

	cmd = exec.Command("sudo", "apt-get", "install", "-y", "nodejs")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running apt-get install nodejs: %v\n", err)
		return
	}

	fmt.Println("Node.js installed successfully using apt.")
}

func checkIfNodeJSIsInstalled(version string) bool {
	cmd := exec.Command("node", "--version")
	cmdOutput, err := cmd.Output()

	fmt.Println("Checking if Node.js is already installed")
	fmt.Println(string(cmdOutput))

	if err == nil && strings.HasPrefix(string(cmdOutput), "v") {
		fmt.Printf("Node.js version %s is already installed. Skipping installation.\n", version)
		return true
	}

  return false
}
