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

func ratingLabel(rating float32) (string, error) {
	if rating < 0 || rating > 10 {
		return "", fmt.Errorf("invalid rating: %.1f", rating)
	}
	if rating >= 7.5 {
		return "great", nil
	}
	if rating >= 5 {
		return "okay", nil
	}
	return "weak", nil
}

func runtimeLabel(mins int) string {
	if mins < 90 {
		return "short"
	}
	if mins >= 90 && mins <= 150 {
		return "standard"
	}

	return "long"
}

func main() {

	title := "Heat"
	year := 1995
	rating := float32(8.3)
	fmt.Println(formatMovie(title, year, rating))

	label, err := ratingLabel(rating)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("verdict:", label)

	fmt.Println()
	fmt.Println(strings.Repeat("-", 15))
	fmt.Println(appInfo())
}
