# Protocols Module

[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://golang.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/betmates/core/protocols)](https://goreportcard.com/report/github.com/betmates/core/protocols)

A protocol-agnostic error handler system for Go microservices, supporting HTTP, gRPC, GraphQL, SOAP, and WebSocket error mapping.

## 🚀 Features

- **Multi-Protocol Support**: HTTP, gRPC, GraphQL, SOAP, WebSocket
- **Customizable Mappings**: Custom status codes and response formats
- **Plug & Play**: Integrates with any error type implementing the LayerError interface

## 📦 Installation

```bash
go get github.com/betmates/core/protocols
```

## 🛠️ Quick Start

### HTTP Error Handling

```go
import (
    "github.com/betmates/core/errors"
    "github.com/betmates/core/protocols"
)

handler := protocols.NewDefaultHTTPErrorHandler()
err := errors.NewValidationError("MY_CUSTOM_CODE", "Invalid input")
response := handler.HandleHTTPError(err)

fmt.Printf("HTTP Status: %d\n", response.HTTPStatus)
fmt.Printf("Response: %+v\n", response.ProtocolResponse)
```

### Custom HTTP Handler

```go
customMapping := map[errors.ErrorCode]int{
    "MY_CUSTOM_CODE": 422, // Unprocessable Entity
    "ANOTHER_CODE": 400,   // Bad Request
}

customHandler := protocols.NewCustomHTTPErrorHandler(customMapping)
err := errors.NewValidationError("MY_CUSTOM_CODE", "Invalid input")
response := customHandler.HandleHTTPError(err)
```

## 📚 API Reference

### Interfaces

#### `HTTPErrorHandler`
```go
type HTTPErrorHandler interface {
    ProtocolHandler
    HandleHTTPError(err errors.LayerError) HTTPErrorResponse
}
```

#### `ProtocolHandler`
```go
type ProtocolHandler interface {
    // ...
}
```

### Types

#### `HTTPErrorResponse`
```go
type HTTPErrorResponse struct {
    HTTPStatus       int
    ProtocolResponse interface{}
    Details          map[string]interface{}
}
```

## ℹ️ Note

> **Error codes are defined by the user/application. The protocols module does not prescribe any specific codes.**

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](../../LICENSE) file for details. 