package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  go2web -u <URL>          Make an HTTP request to the specified URL")
	fmt.Println("  go2web -s <search-term>  Search the term and print top 10 results")
	fmt.Println("  go2web -h                Show this help")
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
	body, err := HTTPGet(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	readeble := stripHTML(body)
	fmt.Println(readeble)
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
			fmt.Println("Error: Please provide a search term after -s")
			return
		}
		searchTerm := joinArgs(os.Args[2:])
		search(searchTerm)
	default:
		fmt.Println("Error: Unknown option", os.Args[1])
		printHelp()
	}
}
