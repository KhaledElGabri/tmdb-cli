package main

import (
	"fmt"
	"strings"
)

const appName = "tmdb-cli"
const version = "0.1"

func appInfo() string {
	return fmt.Sprintf("%s v%s", appName, version)
}

func formatMovie(title string, year int, rating float32) string {
	return fmt.Sprintf("%s (%d) - rating: %.1f", title, year, rating)
}

func main() {

	title := "Heat"
	year := 1995
	rating := float32(8.3)

	fmt.Println(formatMovie(title, year, rating))
	fmt.Println()
	fmt.Println(strings.Repeat("-", 15))

	fmt.Println(appInfo())
}
