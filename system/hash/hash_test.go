package hash

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("HashPassword failed: %v", err)
		return
	}

	// Check if the generated hash is not empty
	if len(hashedPassword) == 0 {
		t.Error("Hashed password is empty")
		return
	}

	// Check if the generated hash is valid
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		t.Errorf("Hashed password is invalid: %v", err)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "password123"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("HashPassword failed: %v", err)
		return
	}

	t.Run("Correct password", func(t *testing.T) {
		// Test with correct password
		if !CheckPasswordHash(password, hashedPassword) {
			t.Error("CheckPasswordHash returned false for correct password")
		}
	})

	t.Run("Incorrect password", func(t *testing.T) {
		// Test with incorrect password
		if CheckPasswordHash("wrongpassword", hashedPassword) {
			t.Error("CheckPasswordHash returned true for incorrect password")
		}
	})
}
