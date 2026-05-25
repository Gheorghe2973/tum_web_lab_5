# Lab 5 — HTTP over TCP Sockets

A command-line HTTP client written in Go that communicates over raw TCP sockets — no built-in or third-party HTTP libraries used.

## Features

- Raw TCP socket communication for both HTTP and HTTPS (via TLS)
- HTTP redirect following (301, 302, 303, 307, 308)
- Chunked transfer encoding (`Transfer-Encoding: chunked`) support
- Custom URL parser (no `net/url` for parsing)
- DuckDuckGo search returning top 10 results
- Human-readable output — HTML tags, scripts, and styles are stripped

## CLI Usage

```
go2web -u <URL>         # make an HTTP request to the specified URL and print the response
go2web -s <search-term> # search the term using DuckDuckGo and print top 10 results
go2web -h               # show this help
```

### Examples

```sh
# Fetch a webpage
go2web -u https://example.com

# Search for something
go2web -s golang tcp sockets

# Open a search result directly
go2web -u https://pkg.go.dev/net

# Show help
go2web -h
```

## How It Works

### HTTP Request Flow

1. The URL is parsed by a hand-written parser ([url.go](lab5/url.go)) that extracts the scheme, host, port, and path.
2. A raw TCP connection is opened via `net.Dial` (HTTP) or `tls.Dial` (HTTPS).
3. An HTTP/1.1 GET request is written directly to the socket.
4. The response is read and the headers are separated from the body.
5. If the response is a redirect, the `Location` header is followed automatically.
6. If `Transfer-Encoding: chunked` is present, the body is decoded chunk by chunk.
7. HTML is stripped from the body so the output is human-readable.

### Search Flow

1. The search term is URL-encoded and appended to the DuckDuckGo HTML endpoint.
2. The HTML response is scanned for result URLs (`uddg=` query params) and their link text.
3. The top 10 unique results are printed with their title and URL.
4. Any result URL can be passed directly to `go2web -u` to open it.

## Project Structure

```
lab5/
├── main.go    # CLI entry point and argument parsing
├── http.go    # Raw TCP HTTP/HTTPS client, redirect and chunked-encoding handling
├── url.go     # Custom URL parser
├── search.go  # DuckDuckGo search and result extraction
└── html.go    # HTML tag stripper and entity decoder
```

## Build & Run

```sh
cd lab5
go build -o go2web .

# Then run:
./go2web -h
./go2web -u https://example.com
./go2web -s open source networking
```

On Windows:

```powershell
cd lab5
go build -o go2web.exe .
.\go2web.exe -u https://example.com
```

## Requirements

- Go 1.21+
- No external dependencies (standard library only)
