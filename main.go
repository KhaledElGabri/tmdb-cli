package main

import (
	"fmt"
	"strings"
)

const appName = "tmdb-cli"
const version = "0.1"

func main() {

	fmt.Println(appName)
	fmt.Println(strings.Repeat("-", 15))
	fmt.Println(version)
}
