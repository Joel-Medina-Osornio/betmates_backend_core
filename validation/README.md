# Validation Module

[![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)](https://golang.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/betting-app/core/validation)](https://goreportcard.com/report/github.com/betting-app/core/validation)

A comprehensive validation system for Go applications, providing functional options-based field validation with extensible rules and localized error messages.

## 🚀 Features

- **🔧 Functional Options API**: Idiomatic Go validation using functional options pattern
- **📝 Extensible Rules**: Built-in validators plus custom validation functions
- **🌍 Localized Messages**: Support for custom error messages in any language
- **⚡ High Performance**: Optimized for production use with minimal overhead
- **🧪 Comprehensive Testing**: Full test coverage with benchmarks

## 📦 Installation

```bash
go get github.com/betting-app/core/validation
```

## 🛠️ Quick Start

### Basic Usage

```go
import (
    "github.com/betting-app/core/validation"
    "github.com/betting-app/core/errors"
)

// Simple validation
err := validation.Validate(
    validation.Field("field1", value1, validation.Required(), validation.MinLength(3)),
    validation.Field("field2", value2, validation.Required(), validation.MaxLength(10)),
)
if err != nil {
    // err is errors.LayerError
    fmt.Printf("Validation error: %s\n", err.Error())
}
```

### Advanced Validation

```go
err := validation.Validate(
    validation.Field("field1", value1, 
        validation.Required("Field1 is required"),
        validation.MinLength(3, "Field1 must have at least 3 characters"),
        validation.MaxLength(20, "Field1 must have at most 20 characters"),
    ),
    validation.Field("field2", value2,
        validation.Required("Field2 is required"),
        validation.MinLength(5, "Field2 must have at least 5 characters"),
    ),
    validation.Field("field3", value3,
        validation.Required("Field3 is required"),
        validation.Pattern(`^[a-zA-Z0-9]+$`, "Field3 must be alphanumeric"),
    ),
    validation.Field("field4", value4,
        validation.Custom(func(value interface{}) bool {
            // Custom validation logic
            return value != nil
        }, "Field4 must not be nil"),
    ),
)
```

## 📚 API Reference

### Core Functions

#### `Validate(fields ...ValidationField) errors.LayerError`
Executes validation on one or more fields and returns the first error found, or nil if all are valid.

#### `Field(field string, value interface{}, opts ...ValidationOption) ValidationField`
Creates a validation field with the specified rules.

### Built-in Validators

#### `Required(msg ...string) ValidationOption`
Validates that the field is not empty (nil, empty string, empty slice/array/map).

#### `MinLength(min int, msg ...string) ValidationOption`
Validates minimum length for strings, slices, and arrays.

#### `MaxLength(max int, msg ...string) ValidationOption`
Validates maximum length for strings, slices, and arrays.

#### `Pattern(pattern string, msg ...string) ValidationOption`
Validates against a regex pattern.

#### `Custom(validator func(value interface{}) bool, msg string) ValidationOption`
Allows custom validation logic.

### Types

```go
type ValidationOption func(field string, value interface{}) errors.LayerError

type ValidationField struct {
    Field   string
    Value   interface{}
    Options []ValidationOption
}
```

## 🔧 Advanced Usage

### Custom Validators

```go
// Custom validator for alphanumeric values
func Alphanumeric(msg ...string) validation.ValidationOption {
    return validation.Pattern(`^[a-zA-Z0-9]+$`, msg...)
}

// Usage
err := validation.Validate(
    validation.Field("field", value, Alphanumeric("Field must be alphanumeric")),
)
```

### Conditional Validation

```go
func ValidateStruct(someStruct SomeStruct) error {
    fields := []validation.ValidationField{
        validation.Field("field1", someStruct.Field1, validation.Required(), validation.MinLength(3)),
    }
    // Only validate field2 if a condition is met
    if someStruct.Condition {
        fields = append(fields, 
            validation.Field("field2", someStruct.Field2, validation.Required(), validation.MaxLength(10)),
        )
    }
    return validation.Validate(fields...)
}
```

## 📊 Performance

### Benchmarks

```bash
go test -bench=. -benchmem ./validation/
```

Example results:
```
BenchmarkValidate_Required-8          1000000    1234 ns/op    512 B/op    8 allocs/op
BenchmarkValidate_MinLength-8         500000     2345 ns/op   1024 B/op   12 allocs/op
BenchmarkValidate_MultipleFields-8    100000    12345 ns/op   2048 B/op   16 allocs/op
```

## 🧪 Testing

### Run Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. -benchmem ./...
```

## ℹ️ Note

> **Error codes and messages are defined by the user/application. The validation module does not prescribe any specific codes or messages.**

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](../../LICENSE) file for details. 