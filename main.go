
 package main

 import (
 	"fmt"
 	"os/exec"
 	"runtime"
 	"strings"
 )

 // Tool represents a command-line tool to check for
 type Tool struct {
 	Name       string
 	Command    string
 	VersionArg string
 }

 // GetOSInfo returns the name of the operating system.
 func GetOSInfo() string {
 	return runtime.GOOS
 }

 // GetToolVersion returns the version of a given tool.
 func GetToolVersion(tool Tool) (string, error) {
 	cmd := exec.Command(tool.Command, tool.VersionArg)
 	output, err := cmd.CombinedOutput()
 	if err != nil {
 		return "", err
 	}
 	return strings.TrimSpace(string(output)), nil
 }

 func main() {
 	fmt.Println("--- System Information ---")

 	// Get and display OS Info
 	osInfo := GetOSInfo()
 	fmt.Printf("Operating System: %s\n", osInfo)

 	fmt.Println("\n--- Tool Versions ---")

 	// Define the tools to check for
 	tools := []Tool{
 		{Name: "Bash", Command: "bash", VersionArg: "--version"},
 		{Name: "Python", Command: "python3", VersionArg: "--version"}, // Prefer python3
 		{Name: "Node.js", Command: "node", VersionArg: "-v"},
 		{Name: "Go", Command: "go", VersionArg: "version"},
 	}

 	for _, tool := range tools {
 		version, err := GetToolVersion(tool)

 		if err != nil {
 			// Try fallback for python
 			if tool.Command == "python3" {
 				py2Tool := Tool{Name: "Python", Command: "python", VersionArg: "--version"}
 				version, err = GetToolVersion(py2Tool)

 				if err != nil {
 					fmt.Printf("%s: Not Found\n", tool.Name)
 					continue
 				}
 			} else {
 				fmt.Printf("%s: Not Found\n", tool.Name)
 				continue
 			}
 		}
 		fmt.Printf("%s: %s\n", tool.Name, version)
 	}
 }
 