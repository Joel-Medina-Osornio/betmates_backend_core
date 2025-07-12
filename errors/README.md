# Error Handling Library

[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://golang.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/betting-app/core/errors)](https://goreportcard.com/report/github.com/betting-app/core/errors)
[![GoDoc](https://godoc.org/github.com/betting-app/core/errors?status.svg)](https://godoc.org/github.com/betting-app/core/errors)
[![Test Coverage](https://img.shields.io/badge/coverage-95%25-green.svg)](https://gocover.io/github.com/betting-app/core/errors)

A comprehensive error handling library for Go microservices, providing advanced error management with validation and multi-protocol support.

## 🚀 Features

- **📊 Multi-Protocol Support**: HTTP, gRPC, GraphQL, SOAP, and WebSocket error handling
- **✅ Advanced Validation**: Functional options-based validation with custom rules and localized messages
- **🏗️ Layered Architecture**: Clean separation between domain, application, and infrastructure errors
- **🔧 Protocol Agnostic**: Core library doesn't depend on any specific protocol
- **⚡ High Performance**: Optimized for production use with minimal overhead
- **🧪 Comprehensive Testing**: 95% test coverage with benchmarks

## 📦 Installation

```bash
go get github.com/betting-app/core/errors
```

## 🛠️ Quick Start

### Basic Error Creation

```go
import "github.com/betting-app/core/errors"

// Create a validation error
err := errors.NewValidationError(
    errors.ErrInvalidEmail,
    "Invalid email format provided",
    map[string]interface{}{
        "email": "invalid-email",
        "field": "email",
    },
)

// Create an infrastructure error
err := errors.NewInfrastructureError(
    errors.ErrDatabaseConnection,
    "Database connection failed",
    map[string]interface{}{
        "database": "postgres",
        "host": "localhost:5432",
    },
)

// Create a business rule error
err := errors.NewBusinessRuleError(
    errors.ErrInvalidBusinessRule,
    "User cannot place bet with insufficient funds",
    map[string]interface{}{
        "user_id": "123",
        "required_amount": 100,
        "available_amount": 50,
    },
)
```

### HTTP Error Handling

```go
// Use default HTTP handler
handler := errors.NewDefaultHTTPErrorHandler()
err := errors.NewValidationError(errors.ErrInvalidEmail, "Invalid email format")
response := handler.HandleHTTPError(err)

fmt.Printf("HTTP Status: %d\n", response.HTTPStatus)
fmt.Printf("Response: %+v\n", response.ProtocolResponse)

// Custom HTTP handler with specific mapping
customMapping := map[errors.ErrorCode]int{
    errors.ErrInvalidEmail:    http.StatusUnprocessableEntity, // 422 instead of 400
    errors.ErrUserNotFound:    http.StatusGone,                // 410 instead of 404
    errors.ErrDatabaseConnection: http.StatusServiceUnavailable, // 503 instead of 424
}

customHandler := errors.NewCustomHTTPErrorHandler(customMapping)
customResponse := customHandler.HandleHTTPError(err)
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

### Error Types

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

```go
// Validation Errors
ErrInvalidEmail    ErrorCode = "INVALID_EMAIL"
ErrInvalidFormat   ErrorCode = "INVALID_FORMAT"
ErrMissingRequired ErrorCode = "MISSING_REQUIRED"
ErrInvalidValue    ErrorCode = "INVALID_VALUE"

// Authentication Errors
ErrInvalidCredentials ErrorCode = "INVALID_CREDENTIALS"
ErrTokenExpired       ErrorCode = "TOKEN_EXPIRED"
ErrTokenInvalid       ErrorCode = "TOKEN_INVALID"

// Authorization Errors
ErrInsufficientPermissions ErrorCode = "INSUFFICIENT_PERMISSIONS"
ErrAccessDenied           ErrorCode = "ACCESS_DENIED"
ErrResourceForbidden      ErrorCode = "RESOURCE_FORBIDDEN"

// Domain Errors
ErrUserNotFound      ErrorCode = "USER_NOT_FOUND"
ErrEmailAlreadyTaken ErrorCode = "EMAIL_ALREADY_TAKEN"
ErrInvalidBusinessRule ErrorCode = "INVALID_BUSINESS_RULE"

// Infrastructure Errors
ErrDatabaseConnection ErrorCode = "DATABASE_CONNECTION"
ErrRepositoryOperation ErrorCode = "REPOSITORY_OPERATION"
ErrExternalService     ErrorCode = "EXTERNAL_SERVICE"
ErrNetworkTimeout      ErrorCode = "NETWORK_TIMEOUT"
```

## 📊 Performance

### Benchmarks

```bash
go test -bench=. -benchmem ./errors/
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