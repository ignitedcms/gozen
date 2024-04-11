package validation

import (
	"testing"
)

func TestValidator_Required(t *testing.T) {
	v := &Validator{}
	v.Required("username", "")
	errors := v.GetErrors()

	if len(errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errors))
	}

	expectedError := ValidationError{Field: "username", Message: "This field is required."}
	if errors[0] != expectedError {
		t.Errorf("Expected error %+v, got %+v", expectedError, errors[0])
	}
}

func TestValidator_MinLength(t *testing.T) {
	v := &Validator{}
	v.MinLength("password", "pass", 6)
	errors := v.GetErrors()

	if len(errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errors))
	}

	expectedError := ValidationError{Field: "password", Message: "This field must be at least 6 characters long."}
	if errors[0] != expectedError {
		t.Errorf("Expected error %+v, got %+v", expectedError, errors[0])
	}
}

func TestValidator_MaxLength(t *testing.T) {
	v := &Validator{}
	v.MaxLength("username", "username", 5)
	errors := v.GetErrors()

	if len(errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errors))
	}

	expectedError := ValidationError{Field: "username", Message: "This field cannot be longer than 5 characters."}
	if errors[0] != expectedError {
		t.Errorf("Expected error %+v, got %+v", expectedError, errors[0])
	}
}

func TestValidator_Email(t *testing.T) {
	v := &Validator{}
	v.Email("email", "invalidemail")
	errors := v.GetErrors()

	if len(errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errors))
	}

	expectedError := ValidationError{Field: "email", Message: "This field must be a valid email address."}
	if errors[0] != expectedError {
		t.Errorf("Expected error %+v, got %+v", expectedError, errors[0])
	}
}

func TestValidator_Min(t *testing.T) {
	v := &Validator{}
	v.Min("age", 17, 18)
	errors := v.GetErrors()

	if len(errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errors))
	}

	expectedError := ValidationError{Field: "age", Message: "This field must be at least 18."}
	if errors[0] != expectedError {
		t.Errorf("Expected error %+v, got %+v", expectedError, errors[0])
	}
}

func TestValidator_Alpha(t *testing.T) {
	v := &Validator{}
	v.Alpha("name", "123")
	errors := v.GetErrors()

	if len(errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errors))
	}

	expectedError := ValidationError{Field: "name", Message: "This field must contain only alphabetic characters."}
	if errors[0] != expectedError {
		t.Errorf("Expected error %+v, got %+v", expectedError, errors[0])
	}
}

func TestValidator_AlphaNum(t *testing.T) {
	v := &Validator{}
	v.AlphaNum("password", "pass@123")
	errors := v.GetErrors()

	if len(errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errors))
	}

	expectedError := ValidationError{Field: "password", Message: "This field must contain only alphanumeric characters."}
	if errors[0] != expectedError {
		t.Errorf("Expected error %+v, got %+v", expectedError, errors[0])
	}
}

func TestValidator_Int(t *testing.T) {
	v := &Validator{}
	v.Int("age", "a12")
	errors := v.GetErrors()

	if len(errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errors))
	}

	expectedError := ValidationError{Field: "age", Message: "This field must be an integer."}
	if errors[0] != expectedError {
		t.Errorf("Expected error %+v, got %+v", expectedError, errors[0])
	}
}

func TestValidator_Float(t *testing.T) {
	v := &Validator{}
	v.Float("weight", "abc")
	errors := v.GetErrors()

	if len(errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errors))
	}

	expectedError := ValidationError{Field: "weight", Message: "This field must be a decimal number."}
	if errors[0] != expectedError {
		t.Errorf("Expected error %+v, got %+v", expectedError, errors[0])
	}
}

func TestValidator_NoErrors(t *testing.T) {
	v := &Validator{}
	if v.HasErrors() {
		t.Error("Expected no errors")
	}
}
