package main

import (
	"fmt"
	"os"
	"strings"
)

const appName = "tmdb-cli"
const version = "0.1"

func appInfo() string {
	return fmt.Sprintf("%s v%s", appName, version)
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("   tmdb popular")
	fmt.Println("   tmdb search <query>")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "popular":
		fmt.Println("Fetching popular movies... (coming soon)")
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("error: search needs a query")
			os.Exit(1)
		}
		query := os.Args[2]
		fmt.Println("Searching for:", query)
	default:
		fmt.Println("unknown command:", command)
		printUsage()
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println(strings.Repeat("-", 15))
	fmt.Println(appInfo())
}
