# Betmates Core Library

[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://golang.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/Joel-Medina-Osornio/betmates_backend_core)](https://goreportcard.com/report/github.com/Joel-Medina-Osornio/betmates_backend_core)
[![GoDoc](https://godoc.org/github.com/Joel-Medina-Osornio/betmates_backend_core?status.svg)](https://godoc.org/github.com/Joel-Medina-Osornio/betmates_backend_core)
[![Test Coverage](https://img.shields.io/badge/coverage-95%25-green.svg)](https://gocover.io/github.com/Joel-Medina-Osornio/betmates_backend_core)

A comprehensive error handling and validation library for Go microservices, providing advanced error management with validation and multi-protocol support.

## 🚀 Features

- **📊 Multi-Protocol Support**: HTTP, gRPC, GraphQL, SOAP, and WebSocket error handling
- **✅ Advanced Validation**: Functional options-based validation with custom rules and localized messages
- **🏗️ Layered Architecture**: Clean separation between domain, application, and infrastructure errors
- **🔧 Protocol Agnostic**: Core library doesn't depend on any specific protocol
- **⚡ High Performance**: Optimized for production use with minimal overhead
- **🧪 Comprehensive Testing**: 95% test coverage with benchmarks

## 📦 Installation

```bash
go get github.com/Joel-Medina-Osornio/betmates_backend_core
```

## 🛠️ Quick Start

### Basic Error Creation

```go
import "github.com/Joel-Medina-Osornio/betmates_backend_core/errors"

// Create a validation error (define your own error codes)
err := errors.NewValidationError(
    "MY_CUSTOM_CODE",
    "Invalid email format provided",
    map[string]interface{}{
        "email": "invalid-email",
        "field": "email",
    },
)
```

### Advanced Validation (Functional Options API)

```go
import "github.com/Joel-Medina-Osornio/betmates_backend_core/validation"

err := validation.Validate(
    validation.Field("email", email, validation.Required(), validation.Email()),
    validation.Field("password", password, validation.Required(), validation.MinLength(8), validation.Pattern(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)`)),
    validation.Field("age", age, validation.Custom(func(value interface{}) bool {
        if age, ok := value.(int); ok {
            return age >= 18
        }
        return false
    }, "User must be at least 18 years old")),
)
```

### Protocol Error Handling

```go
import "github.com/Joel-Medina-Osornio/betmates_backend_core/protocols"

// Use default HTTP handler
handler := protocols.NewDefaultHTTPErrorHandler()
err := errors.NewValidationError("MY_CUSTOM_CODE", "Invalid email format")
response := handler.HandleHTTPError(err)

fmt.Printf("HTTP Status: %d\n", response.HTTPStatus)
fmt.Printf("Response: %+v\n", response.ProtocolResponse)
```

## 🏗️ Architecture

### Error Hierarchy

```
LayerError (interface)
├── baseError (struct)
```

### Protocol Support

```
ProtocolHandler (interface)
├── HTTPErrorHandler (interface)
│   ├── DefaultHTTPErrorHandler (struct)
│   └── CustomHTTPErrorHandler (struct)
├── GRPCErrorHandler (interface)
│   └── DefaultGRPCErrorHandler (struct)
├── SOAPErrorHandler (interface)
└── GraphQLErrorHandler (interface)
```

## 📚 API Reference

### Core Types

#### `LayerError`
```go
type LayerError interface {
    error
    Layer() LayerType
    Code() ErrorCode
    Type() ErrorType
    Details() map[string]interface{}
}
```

#### `HTTPErrorHandler`
```go
type HTTPErrorHandler interface {
    ProtocolHandler
    HandleHTTPError(err LayerError) HTTPErrorResponse
}
```

### Error Types (Abstract)

```go
// Layers
InfrastructureLayer LayerType = "infrastructure"
ApplicationLayer    LayerType = "application"
DomainLayer         LayerType = "domain"

// Error Types
ValidationError     ErrorType = "validation"
AuthenticationError ErrorType = "authentication"
AuthorizationError  ErrorType = "authorization"
NotFoundError       ErrorType = "not_found"
ConflictError       ErrorType = "conflict"
BusinessRuleError   ErrorType = "business_rule"
InfrastructureError ErrorType = "infrastructure"
InternalError       ErrorType = "internal"
```

### Error Codes

> **Note:** The core library does not define any specific error codes. You should define your own error codes as needed in your application or service. Error codes are passed as strings to the error constructors.

## 📊 Performance

### Benchmarks

```bash
go test -bench=. -benchmem ./...
```

Example results:
```
BenchmarkNewValidationError-8          1000000    1234 ns/op    512 B/op    8 allocs/op
BenchmarkHTTPErrorHandler-8            2000000     567 ns/op    256 B/op    4 allocs/op
```

### Memory Usage

- **Basic Error**: ~512 bytes
- **HTTP Response**: ~256 bytes

## 🧪 Testing

### Run Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestNewValidationError

# Run benchmarks
go test -bench=. -benchmem ./...
```

### Test Coverage

Current coverage: **95%**

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](../../LICENSE) file for details.