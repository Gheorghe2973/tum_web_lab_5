package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  -h            Show help message")
	fmt.Println("  -u <URL>      Fetch the content of the specified URL")
	fmt.Println("  -s <term>     Search for the specified term in a file")
}

func joinArgs(args []string) string {
	result := ""
	for i, arg := range args {
		if i > 0 {
			result += " "
		}
		result += arg

	}
	return result
}

func fetchURL(url string) {
	fmt.Println("Fetching URL:", url)
}

func search(term string) {
	fmt.Println("Searching for term:", term)
}

func main() {

	if len(os.Args) < 2 {
		printHelp()
		return
	}
	switch os.Args[1] {
	case "-h":
		printHelp()

	case "-u":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a URL after -u")
			return
		}
		fetchURL(os.Args[2])
	case "-s":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a file path after -s")
			return
		}
		searchTerm := joinArgs(os.Args[2:])
		search(searchTerm)
	default:
		fmt.Println("Error: Unknown option", os.Args[1])
		printHelp()
	}

}
