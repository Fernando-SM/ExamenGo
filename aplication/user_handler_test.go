package aplication

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserRegisterHandler(t *testing.T) {
	// Simular la solicitud de registro de usuario.
	user := User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		Phone:    "1234567890",
	}
	userJSON, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UserRegisterHandler)

	// Ejecutar el manejador.
	handler.ServeHTTP(rr, req)

	// Chequear el código de estado esperado.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Chequear la respuesta del cuerpo.
	expected := `{"token":"someTokenValue"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUserLoginHandler(t *testing.T) {
	// Simular la solicitud de inicio de sesión de usuario.
	var creds = struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    "test@example.com",
		Password: "password123",
	}
	credsJSON, _ := json.Marshal(creds)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(credsJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UserLoginHandler)

	// Ejecutar el manejador.
	handler.ServeHTTP(rr, req)

	// Chequear el código de estado esperado.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Chequear la respuesta del cuerpo.
	expected := `{"token":"someTokenValue"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
