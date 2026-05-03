package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"strings"
)

func HTTPGet(rawURL string) (string, error) {
	parsed, err := parseUrl(rawURL)

	if err != nil {
		return "", err
	}

	address := parsed.Host + ":" + parsed.Port
	var conn net.Conn

	if parsed.Scheme == "https" {
		conn, err = tls.Dial("tcp", address, &tls.Config{})
	} else {
		conn, err = net.Dial("tcp", address)
	}
	if err != nil {
		return "", fmt.Errorf("connection failed: %w", err)
	}
	defer conn.Close()

	request := fmt.Sprintf(
		"GET %s HTTP/1.1\r\nHost: %s\r\nConnection: close\r\nAccept: text/html\r\nAccept-Encoding: identity\r\nUser-Agent: go2web/1.0\r\n\r\n",
		parsed.Path, parsed.Host)

	_, err = conn.Write([]byte(request))

	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	reader := bufio.NewReader(conn)

	responseBytes, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	response := string(responseBytes)

	headerEnd := strings.Index(response, "\r\n\r\n")

	if headerEnd == -1 {
		return "", fmt.Errorf("invalid HTTP response: no header-body separator found")
	}

	headers := response[:headerEnd]
	body := response[headerEnd+4:]

	statusLine := strings.SplitN(headers, "\r\n", 2)[0]

	if strings.Contains(statusLine, "301") || strings.Contains(statusLine, "302") || strings.Contains(statusLine, "303") || strings.Contains(statusLine, "307") || strings.Contains(statusLine, "308") {
		location := getHeader(headers, "Location")
		if location != "" {
			fmt.Println("Redirecting to:", location)
			if strings.HasPrefix(location, "/") {
				location = parsed.Scheme + "://" + parsed.Host + location
			}
			return HTTPGet(location)
		}
	}
	if strings.Contains(strings.ToLower(headers), "transfer-encoding: chunked") {
		body = decodeChunked(body)
	}

	return body, nil

}
func getHeader(headers string, name string) string {
	lines := strings.Split(headers, "\r\n")
	nameLower := strings.ToLower(name)
	for _, line := range lines {
		if strings.HasPrefix(strings.ToLower(line), nameLower+":") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}

func decodeChunked(body string) string {
	var result strings.Builder
	remaining := body

	for {
		idx := strings.Index(remaining, "\r\n")
		if idx == -1 {
			break
		}
		sizeHex := strings.TrimSpace(remaining[:idx])
		if sizeHex == "" || sizeHex == "0" {
			break
		}
		var size int
		fmt.Sscanf(sizeHex, "%x", &size)
		if size == 0 {
			break
		}
		dataStart := idx + 2
		if dataStart+size > len(remaining) {
			break
		}
		result.WriteString(remaining[dataStart : dataStart+size])
		remaining = remaining[dataStart+size+2:]
	}
	return result.String()
}
