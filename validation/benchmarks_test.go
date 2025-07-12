package validation

import (
	"testing"
)

// Benchmark: Validar un email inválido
func BenchmarkValidateEmail(b *testing.B) {
	email := "invalid-email"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Validate(Field("email", email, Email()))
	}
}

// Benchmark: Validar un password simple
func BenchmarkValidatePassword(b *testing.B) {
	password := "weak"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Validate(Field("password", password, Required(), MinLength(8)))
	}
}

// Benchmark: Validar múltiples campos comunes
func BenchmarkValidateMultipleFields(b *testing.B) {
	email := "invalid-email"
	password := "weak"
	age := 15
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Validate(
			Field("email", email, Required(), Email()),
			Field("password", password, Required(), MinLength(8)),
			Field("age", age, Custom(func(value interface{}) bool {
				if age, ok := value.(int); ok {
					return age >= 18
				}
				return false
			}, "User must be at least 18 years old")),
		)
	}
}
