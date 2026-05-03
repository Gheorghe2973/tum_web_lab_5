package main

import (
	"fmt"
	"net/url"
	"strings"
)

type SearchResult struct {
	Title string
	URL   string
}

func search(term string) {
	encode := url.QueryEscape(term)

	searchURL := "https://html.duckduckgo.com/html/?q=" + encode

	body, err := HTTPGet(searchURL)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	results := extractSearchResults(body)

	if len(results) == 0 {
		fmt.Println("No results found.")
		return
	}

	count := 10
	if len(results) < count {
		count = len(results)
	}

	for i := 0; i < count; i++ {
		println(i+1, ".", results[i].Title)
		println("   ", results[i].URL)
		println()
	}
}

func extractSearchResults(html string) []SearchResult {
	var results []SearchResult

	remaining := html

	for {
		idx := strings.Index(remaining, "uddg=")
		if idx == -1 {
			break
		}
		remaining = remaining[idx+len("uddg="):]

		endIdx := strings.Index(remaining, "&")
		if endIdx == -1 {
			break
		}
		resultURL := remaining[:endIdx]
		decodedURL, _ := url.QueryUnescape(resultURL)

		if !strings.HasPrefix(decodedURL, "http") {
			continue
		}

		dublicate := false
		for _, r := range results {
			if r.URL == decodedURL {
				dublicate = true
				break
			}
		}
		if dublicate {
			continue
		}
		title := extractTitle(remaining)
		if title == "" {
			title = decodedURL
		}
		results = append(results, SearchResult{
			Title: title,
			URL:   decodedURL,
		})
	}
	return results
}

func extractTitle(html string) string {
	aStart := strings.Index(html, ">")
	if aStart == -1 {
		return ""
	}
	contentStart := aStart + 1

	aEnd := strings.Index(html[contentStart:], "</a>")
	if aEnd == -1 {
		return ""
	}
	titleHTML := html[contentStart : contentStart+aEnd]
	return strings.TrimSpace(stripHTML(titleHTML))
}
