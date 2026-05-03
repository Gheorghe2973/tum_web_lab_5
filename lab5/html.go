package main

import (
	"strings"
)

func strimHTML(html string) string {
	var result strings.Builder
	inTag := false

	// Remove <style>...</style> blocks
	for {
		start := strings.Index(strings.ToLower(html), "<style")
		if start == -1 {
			break
		}
		end := strings.Index(strings.ToLower(html[start:]), "</style>")
		if end == -1 {
			break
		}
		html = html[:start] + html[start+end+len("</style>"):]
	}

	// Remove <script>...</script> blocks
	for {
		start := strings.Index(strings.ToLower(html), "<script")
		if start == -1 {
			break
		}
		end := strings.Index(strings.ToLower(html[start:]), "</script>")
		if end == -1 {
			break
		}
		html = html[:start] + html[start+end+len("</script>"):]
	}

	for _, ch := range html {
		switch {
		case ch == '<':
			inTag = true
		case ch == '>':
			inTag = false
			result.WriteRune(' ')
		case !inTag:
			result.WriteRune(ch)
		}
	}
	text := result.String()

	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&quot;", "\"")
	text = strings.ReplaceAll(text, "&apos;", "'")
	text = strings.ReplaceAll(text, "&#39;", "'")
	text = strings.ReplaceAll(text, "&nbsp;", " ")

	lines := strings.Split(text, "\n")
	var cleaned []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			cleaned = append(cleaned, line)
		}
	}
	return strings.Join(cleaned, "\n")
}
