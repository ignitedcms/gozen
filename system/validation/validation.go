/*
|---------------------------------------------------------------
| Validation utility helper
|---------------------------------------------------------------
|
| Custom validation helper with no third party dependencies
| Work needs to be done on email and date validation
|
| @license: MIT
| @version: 1.0
| @since: 1.0
*/
package validation

import (
	"fmt"
	"gozen/db"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ValidationError struct {
	Field   string
	Message string
}

type Validator struct {
	errors []ValidationError
}

func (v *Validator) Unique(field, value, table, column string) *Validator {
	// Check if the value is unique in the given table and column
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", table, column)
	var count int
	err := db.DB.QueryRow(query, value).Scan(&count)
	if err != nil {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "Error checking for uniqueness: " + err.Error()})
		return v
	}

	if count > 0 {
		v.errors = append(v.errors, ValidationError{Field: field, Message: fmt.Sprintf("The value '%s' is not unique in the '%s' table", value, table)})
	}

	return v
}

func (v *Validator) Exists(field, value, table, column string) *Validator {
	// Check if the value exists in the given table and column
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", table, column)
	var count int
	err := db.DB.QueryRow(query, value).Scan(&count)
	if err != nil {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "Error checking for existence: " + err.Error()})
		return v
	}

	if count == 0 {
		v.errors = append(v.errors, ValidationError{Field: field, Message: fmt.Sprintf("The value '%s' does not exist in the '%s' table ", value, table)})
	}

	return v
}

func (v *Validator) Required(field, value string) *Validator {
	if strings.TrimSpace(value) == "" {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "This field is required."})
	}
	return v
}

func (v *Validator) MinLength(field, value string, minLength int) *Validator {
	if len(value) < minLength {
		v.errors = append(v.errors, ValidationError{Field: field, Message: fmt.Sprintf("This field must be at least %d characters long.", minLength)})
	}
	return v
}

func (v *Validator) MaxLength(field, value string, maxLength int) *Validator {
	if len(value) > maxLength {
		v.errors = append(v.errors, ValidationError{Field: field, Message: fmt.Sprintf("This field cannot be longer than %d characters.", maxLength)})
	}
	return v
}

// This broken use mail.ParseAddress(email)
func (v *Validator) Email(field, value string) *Validator {
	_, err := mail.ParseAddress(value)
	if err != nil {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "This field must be a valid email address."})
	}
	return v
}

func (v *Validator) Min(field string, value int, min int) *Validator {
	if value < min {
		v.errors = append(v.errors, ValidationError{Field: field, Message: fmt.Sprintf("This field must be at least %d.", min)})
	}
	return v
}

func (v *Validator) Alpha(field, value string) *Validator {
	alphaRegex := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !alphaRegex.MatchString(value) {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "This field must contain only alphabetic characters."})
	}
	return v
}

func (v *Validator) AlphaNum(field, value string) *Validator {
	alphaNumRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !alphaNumRegex.MatchString(value) {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "This field must contain only alphanumeric characters."})
	}
	return v
}

func (v *Validator) Int(field string, value string) *Validator {
	_, err := strconv.Atoi(value)
	if err != nil {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "This field must be an integer."})
	}
	return v
}

func (v *Validator) Float(field string, value string) *Validator {
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "This field must be a decimal number."})
	}
	return v
}

func (v *Validator) Date(field, value string) *Validator {
	_, err := time.Parse("2006-01-02", value)
	if err != nil {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "This field must be a valid date in the format YYYY-MM-DD."})
	}
	return v
}

func (v *Validator) DateLessThan(field, value, maxDate string) *Validator {
	date, err := time.Parse("2006-01-02", value)
	if err != nil {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "This field must be a valid date in the format YYYY-MM-DD."})
		return v
	}

	max, err := time.Parse("2006-01-02", maxDate)
	if err != nil {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "Invalid maximum date format. Expected YYYY-MM-DD."})
		return v
	}

	if date.After(max) {
		v.errors = append(v.errors, ValidationError{Field: field, Message: fmt.Sprintf("This date must be less than %s.", maxDate)})
	}

	return v
}

func (v *Validator) DateGreaterThan(field, value, minDate string) *Validator {
	date, err := time.Parse("2006-01-02", value)
	if err != nil {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "This field must be a valid date in the format YYYY-MM-DD."})
		return v
	}

	min, err := time.Parse("2006-01-02", minDate)
	if err != nil {
		v.errors = append(v.errors, ValidationError{Field: field, Message: "Invalid minimum date format. Expected YYYY-MM-DD."})
		return v
	}

	if date.Before(min) {
		v.errors = append(v.errors, ValidationError{Field: field, Message: fmt.Sprintf("This date must be greater than %s.", minDate)})
	}

	return v
}

func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

func (v *Validator) GetErrors() []ValidationError {
	return v.errors
}

/*
Usage----
formData := FormData{
   Name:     r.FormValue("name"),
   Email:    r.FormValue("email"),
   Age:      0, // You might need to parse the age from string to int
   Salary:   0, // You might need to parse the salary from string to float64
   Username: r.FormValue("username"),
}

validator := &validation.Validator{}

validator.Required("Name", formData.Name).
MinLength("Name", formData.Name, 3).
MaxLength("Name", formData.Name, 50).
Required("Email", formData.Email).
Email("Email", formData.Email).
Min("Age", formData.Age, 18).
Required("Salary", fmt.Sprintf("%.2f", formData.Salary)).
Float("Salary", fmt.Sprintf("%.2f", formData.Salary)).
Required("Username", formData.Username).
AlphaNum("Username", formData.Username)

if validator.HasErrors() {
   // Handle validation errors
   for _, err := range validator.GetErrors() {
      fmt.Fprintf(w, "%s: %s\n", err.Field, err.Message)
   }
   return
}
*/
