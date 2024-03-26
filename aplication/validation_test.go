package aplication

import (
	"testing"
)

// TestValidateEmail verifica que la función ValidateEmail trabaje como se espera.
func TestValidateEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"invalid-email", false},
		// Añade más casos de prueba aquí...
	}

	for _, test := range tests {
		if result := ValidateEmail(test.email); result != test.expected {
			t.Errorf("ValidateEmail(%v) = %v; want %v", test.email, result, test.expected)
		}
	}
}

// TestValidatePhone verifica que la función ValidatePhone trabaje como se espera.
func TestValidatePhone(t *testing.T) {
	tests := []struct {
		phone    string
		expected bool
	}{
		{"1234567890", true},
		{"123", false},
		// Añade más casos de prueba aquí...
	}

	for _, test := range tests {
		if result := ValidatePhone(test.phone); result != test.expected {
			t.Errorf("ValidatePhone(%v) = %v; want %v", test.phone, result, test.expected)
		}
	}
}

// TestValidatePassword verifica que la función ValidatePassword trabaje como se espera.
func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"Passw0rd@", true},
		{"short", false},
		// Añade más casos de prueba aquí...
	}

	for _, test := range tests {
		if result := ValidatePassword(test.password); result != test.expected {
			t.Errorf("ValidatePassword(%v) = %v; want %v", test.password, result, test.expected)
		}
	}
}

// TestValidateUserRegistration verifica que la función ValidateUserRegistration funcione correctamente.
func TestValidateUserRegistration(t *testing.T) {
	user := user.User{
		Username: "validUser",
		Email:    "test@example.com",
		Phone:    "1234567890",
		Password: "Passw0rd@",
	}

	// Caso exitoso
	if err := ValidateUserRegistration(user); err != nil {
		t.Errorf("ValidateUserRegistration() with valid data should not return an error, got: %v", err)
	}

	// Caso con email no válido
	user.Email = "bademail"
	if err := ValidateUserRegistration(user); err == nil {
		t.Error("ValidateUserRegistration() with invalid email should return an error")
	}
}
