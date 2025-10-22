package main

import (
	"errne/errne"
	"fmt"
)

func main() {
	fmt.Println("--- System Information ---")

	// Get and display OS Info
	osInfo := errne.GetOSInfo()
	fmt.Printf("Operating System: %s\n", osInfo)

	fmt.Println("\n--- Tool Versions ---")

	// Define the tools to check for

tools := []errne.Tool{
		{Name: "Bash", Command: "bash", VersionArg: "--version"},
		{Name: "Python", Command: "python3", VersionArg: "--version"}, // Prefer python3
		{Name: "Node.js", Command: "node", VersionArg: "-v"},
		{Name: "Go", Command: "go", VersionArg: "version"},
	}

	for _, tool := range tools {
		version, err := errne.GetToolVersion(tool)

		if err != nil {
			// Try fallback for python
			if tool.Command == "python3" {
				py2Tool := errne.Tool{Name: "Python", Command: "python", VersionArg: "--version"}
				version, err = errne.GetToolVersion(py2Tool)

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