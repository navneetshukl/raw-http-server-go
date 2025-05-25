# HTTP Web Server in Go

A lightweight HTTP web server implementation built from scratch in Go using only the standard library. This project is part of the [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-webserver/) series.

## Features

- **Pure Go Implementation**: Built using only Go's standard library without external dependencies
- **Concurrent Request Handling**: Supports multiple simultaneous connections using goroutines
- **HTTP/1.1 Protocol**: Implements basic HTTP/1.1 request parsing and response generation
- **Static File Serving**: Serves HTML files for root and index paths
- **Configurable Port**: Port configuration via environment variable
- **Request Logging**: Detailed logging of incoming HTTP requests

## Getting Started

### Prerequisites

- Go 1.16 or higher installed on your system

### Installation

1. Clone the repository:
```bash
git clone <your-repository-url>
cd http-webserver
```

2. Build the application:
```bash
go build -o webserver main.go
```

### Running the Server

#### Default Port (8080)
```bash
./webserver
```

#### Custom Port
```bash
port=3000 ./webserver
```

Or set the environment variable:
```bash
export port=3000
./webserver
```

## Usage

Once the server is running, you can access it via:

- **Root Path**: `http://localhost:8080/` - Serves HTML content
- **Index Path**: `http://localhost:8080/index` - Serves HTML content  
- **Other Paths**: `http://localhost:8080/any-path` - Returns plain text with the requested path

### Example Requests

```bash
# Test basic functionality
curl http://localhost:8080/

# Test custom paths
curl http://localhost:8080/api/users
curl http://localhost:8080/about
```

## Implementation Details

### Architecture

The server implements a simple HTTP server with the following components:

1. **Main Server Loop**: Listens for incoming TCP connections
2. **Connection Handler**: Processes each connection in a separate goroutine
3. **Request Parser**: Parses HTTP requests and extracts path information
4. **Response Generator**: Creates appropriate HTTP responses

### Key Functions

- `main()`: Sets up the server listener and accepts incoming connections
- `handleMultipleConnection()`: Processes individual HTTP requests and generates responses
- `loadHTML()`: Loads and serves HTML content (function referenced but not shown in provided code)

### HTTP Request Handling

The server handles HTTP requests by:

1. Reading raw request data from the TCP connection
2. Parsing the request line to extract the HTTP method and path
3. Validating that the request is a valid GET request
4. Generating appropriate responses based on the requested path

### Response Types

- **HTML Response**: For `/` and `/index` paths
  - Content-Type: `text/html`
  - Serves content from HTML file
  
- **Plain Text Response**: For all other paths
  - Content-Type: `text/plain`
  - Returns the requested path information

## Project Structure

```
.
├── main.go          # Main server implementation
├── README.md        # This file
└── index.html       # HTML file served for root/index paths (if applicable)
```

## Error Handling

The server includes error handling for:

- Port binding failures
- Connection acceptance errors
- Request reading errors
- File loading errors
- Response writing errors

## Limitations

This is a basic HTTP server implementation for educational purposes and has several limitations:

- Only supports GET requests
- No support for HTTP headers processing
- No support for query parameters
- No support for request body parsing
- Basic error handling
- No HTTPS support
- No authentication or authorization

## Contributing

This project is part of a coding challenge. Feel free to fork and experiment with additional features such as:

- POST request support
- Query parameter parsing
- Better error handling
- Static file serving from directories
- Basic routing
- Middleware support

## License

This project is open source and available under the [MIT License](LICENSE).

## Acknowledgments

- [Coding Challenges](https://codingchallenges.fyi/) for the problem statement
- Go standard library documentation

## Learning Resources

- [Go net package documentation](https://pkg.go.dev/net)
- [HTTP/1.1 Specification (RFC 7230)](https://tools.ietf.org/html/rfc7230)
- [Building Web Servers in Go](https://golang.org/doc/articles/wiki/)