package validation

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/Joel-Medina-Osornio/betmates_backend_core/errors"
)

type ValidationOption func(field string, value any) errors.LayerError

func Field(field string, value any, opts ...ValidationOption) ValidationField {
	return ValidationField{field, value, opts}
}

type ValidationField struct {
	Field   string
	Value   any
	Options []ValidationOption
}

func Validate(fields ...ValidationField) errors.LayerError {
	for _, f := range fields {
		for _, opt := range f.Options {
			if err := opt(f.Field, f.Value); err != nil {
				return err
			}
		}
	}
	return nil
}

func Required(msg ...string) ValidationOption {
	return func(field string, value any) errors.LayerError {
		message := fmt.Sprintf("Field '%s' is required", field)
		if len(msg) > 0 {
			message = msg[0]
		}
		if value == nil {
			return errors.NewValidationError(errors.ErrMissingRequired, message, map[string]any{"field": field})
		}
		val := reflect.ValueOf(value)
		switch val.Kind() {
		case reflect.String:
			if strings.TrimSpace(val.String()) == "" {
				return errors.NewValidationError(errors.ErrMissingRequired, message, map[string]any{"field": field})
			}
		case reflect.Slice, reflect.Array, reflect.Map:
			if val.Len() == 0 {
				return errors.NewValidationError(errors.ErrMissingRequired, message, map[string]any{"field": field})
			}
		}
		return nil
	}
}

func Email(msg ...string) ValidationOption {
	return func(field string, value any) errors.LayerError {
		message := fmt.Sprintf("Field '%s' must be a valid email", field)
		if len(msg) > 0 {
			message = msg[0]
		}
		email, ok := value.(string)
		if !ok || email == "" {
			return errors.NewValidationError(errors.ErrInvalidEmail, message, map[string]any{"field": field})
		}
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(email) {
			return errors.NewValidationError(errors.ErrInvalidEmail, message, map[string]any{"field": field, "value": email})
		}
		return nil
	}
}

func MinLength(min int, msg ...string) ValidationOption {
	return func(field string, value any) errors.LayerError {
		message := fmt.Sprintf("Field '%s' must have at least %d characters", field, min)
		if len(msg) > 0 {
			message = msg[0]
		}
		val := reflect.ValueOf(value)
		switch val.Kind() {
		case reflect.String:
			if len(val.String()) < min {
				return errors.NewValidationError(errors.ErrInvalidFormat, message, map[string]any{"field": field, "min": min})
			}
		case reflect.Slice, reflect.Array:
			if val.Len() < min {
				return errors.NewValidationError(errors.ErrInvalidFormat, message, map[string]any{"field": field, "min": min})
			}
		}
		return nil
	}
}

func MaxLength(max int, msg ...string) ValidationOption {
	return func(field string, value any) errors.LayerError {
		message := fmt.Sprintf("Field '%s' must have at most %d characters", field, max)
		if len(msg) > 0 {
			message = msg[0]
		}
		val := reflect.ValueOf(value)
		switch val.Kind() {
		case reflect.String:
			if len(val.String()) > max {
				return errors.NewValidationError(errors.ErrInvalidFormat, message, map[string]any{"field": field, "max": max})
			}
		case reflect.Slice, reflect.Array:
			if val.Len() > max {
				return errors.NewValidationError(errors.ErrInvalidFormat, message, map[string]any{"field": field, "max": max})
			}
		}
		return nil
	}
}

func Pattern(pattern string, msg ...string) ValidationOption {
	return func(field string, value any) errors.LayerError {
		message := fmt.Sprintf("Field '%s' must match pattern", field)
		if len(msg) > 0 {
			message = msg[0]
		}
		str, ok := value.(string)
		if !ok {
			return errors.NewValidationError(errors.ErrInvalidFormat, message, map[string]any{"field": field})
		}
		regex := regexp.MustCompile(pattern)
		if !regex.MatchString(str) {
			return errors.NewValidationError(errors.ErrInvalidFormat, message, map[string]any{"field": field, "pattern": pattern})
		}
		return nil
	}
}

func Custom(validator func(value any) bool, msg string) ValidationOption {
	return func(field string, value any) errors.LayerError {
		if !validator(value) {
			return errors.NewValidationError(errors.ErrInvalidFormat, msg, map[string]any{"field": field})
		}
		return nil
	}
}
