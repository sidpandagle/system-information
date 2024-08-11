package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// getHostname fetches the hostname of the system
func getHostname() (string, error) {
	return os.Hostname()
}

// getOSInfo fetches the operating system and architecture
func getOSInfo() (string, string) {
	return runtime.GOOS, runtime.GOARCH
}

// getUptime fetches the system uptime
func getUptime() string {
	uptimeCmd := exec.Command("uptime", "-p")
	var out bytes.Buffer
	uptimeCmd.Stdout = &out
	err := uptimeCmd.Run()
	if err != nil {
		return err.Error()
	}
	return out.String()
}

// getCPUInfo fetches CPU details
func getCPUInfo() string {
	cmd := exec.Command("lscpu")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err.Error()
	}
	return out.String()
}

func main() {
	// OS and architecture
	osName, arch := getOSInfo()
	fmt.Printf("Operating System: %s\nArchitecture: %s\n", osName, arch)

	// Hostname
	hostname, err := getHostname()
	if err != nil {
		fmt.Println("Error fetching hostname:", err)
	} else {
		fmt.Printf("Hostname: %s\n", hostname)
	}

	// Uptime
	uptime := getUptime()
	fmt.Printf("Uptime: %s\n", uptime)

	// CPU Info
	cpuInfo := getCPUInfo()
	fmt.Printf("CPU Info:\n%s\n", cpuInfo)

	// Current time
	fmt.Printf("Current Time: %s\n", time.Now().Format(time.RFC1123))
}
