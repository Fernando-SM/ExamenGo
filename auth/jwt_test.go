package auth

import (
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	jwtManager := NewJWTManager("testkey")
	username := "testuser"
	token, err := jwtManager.GenerateJWT(username, 1*time.Hour)
	if err != nil {
		t.Errorf("Failed to generate JWT: %v", err)
	}
	if token == "" {
		t.Errorf("Generated JWT is empty")
	}
	// Aquí se agregarían más pruebas para comprobar la validez del token, etc.
}
