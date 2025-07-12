package validation

import (
	"testing"
)

func TestValidate_Required(t *testing.T) {
	err := Validate(Field("name", "", Required()))
	if err == nil {
		t.Error("Expected error for empty required field")
	}
	err = Validate(Field("name", "John", Required()))
	if err != nil {
		t.Errorf("Did not expect error for non-empty required field: %v", err)
	}
}

func TestValidate_Email(t *testing.T) {
	err := Validate(Field("email", "invalid-email", Email()))
	if err == nil {
		t.Error("Expected error for invalid email")
	}
	err = Validate(Field("email", "test@example.com", Email()))
	if err != nil {
		t.Errorf("Did not expect error for valid email: %v", err)
	}
}

func TestValidate_MinLength(t *testing.T) {
	err := Validate(Field("password", "123", MinLength(8)))
	if err == nil {
		t.Error("Expected error for short password")
	}
	err = Validate(Field("password", "12345678", MinLength(8)))
	if err != nil {
		t.Errorf("Did not expect error for valid length: %v", err)
	}
}

func TestValidate_MaxLength(t *testing.T) {
	err := Validate(Field("username", "verylongusername", MaxLength(10)))
	if err == nil {
		t.Error("Expected error for long username")
	}
	err = Validate(Field("username", "short", MaxLength(10)))
	if err != nil {
		t.Errorf("Did not expect error for valid length: %v", err)
	}
}

func TestValidate_Pattern(t *testing.T) {
	err := Validate(Field("phone", "123-456-7890", Pattern(`^\d{3}-\d{3}-\d{4}$`)))
	if err != nil {
		t.Errorf("Did not expect error for valid pattern: %v", err)
	}
	err = Validate(Field("phone", "1234567890", Pattern(`^\d{3}-\d{3}-\d{4}$`)))
	if err == nil {
		t.Error("Expected error for invalid pattern")
	}
}

func TestValidate_Custom(t *testing.T) {
	isAdult := func(value interface{}) bool {
		if age, ok := value.(int); ok {
			return age >= 18
		}
		return false
	}
	err := Validate(Field("age", 15, Custom(isAdult, "Must be at least 18 years old")))
	if err == nil {
		t.Error("Expected error for age < 18")
	}
	err = Validate(Field("age", 20, Custom(isAdult, "Must be at least 18 years old")))
	if err != nil {
		t.Errorf("Did not expect error for age >= 18: %v", err)
	}
}

func TestValidate_MultipleFields(t *testing.T) {
	err := Validate(
		Field("email", "invalid", Required(), Email()),
		Field("password", "123", Required(), MinLength(8)),
	)
	if err == nil {
		t.Error("Expected error for multiple invalid fields")
	}
}
