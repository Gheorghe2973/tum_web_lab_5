package main

import (
	"strings"
)

type ParsedUrl struct {
	Scheme string
	Host   string
	Port   string
	Path   string
}

func parseUrl(rawUrl string) (ParsedUrl, error) {
	result := ParsedUrl{}

	if strings.HasPrefix(rawUrl, "https://") {
		result.Scheme = "https"
		result.Port = "443"
		rawUrl = rawUrl[len("https://"):]
	} else if strings.HasPrefix(rawUrl, "http://") {
		result.Scheme = "http"
		result.Port = "80"
		rawUrl = rawUrl[len("http://"):]
	} else {
		result.Scheme = "https"
		result.Port = "443"
	}

	parts := strings.SplitN(rawUrl, "/", 2)
	hostPart := parts[0]
	if len(parts) > 1 {
		result.Path = "/" + parts[1]
	} else {
		result.Path = "/"
	}

	if strings.Contains(hostPart, ":") {
		hostParts := strings.SplitN(hostPart, ":", 2)
		result.Host = hostParts[0]
		result.Port = hostParts[1]
	} else {
		result.Host = hostPart
	}
	return result, nil
}
