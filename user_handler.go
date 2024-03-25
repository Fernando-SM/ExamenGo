package main

import (
	"encoding/json"
	"net/http"
)

// UserRegisterHandler maneja la solicitud de registro de usuarios.
func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decodifica el cuerpo de la solicitud JSON en la estructura User.
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Valida los datos del usuario (email, teléfono, contraseña).
	if err := ValidateUserRegistration(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Agrega el usuario a la "base de datos" simulada (usersDB).
	usersDB = append(usersDB, user)

	// Genera un token JWT para el usuario recién registrado.
	token, err := GenerateJWT(user)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Envía una respuesta de éxito con el token JWT.
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// UserLoginHandler maneja la solicitud de inicio de sesión de usuarios.
func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decodifica el cuerpo de la solicitud JSON en la estructura creds.
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Verifica si los campos email y password están presentes en la solicitud.
	if creds.Email == "" {
		http.Error(w, "Falta el campo correo", http.StatusBadRequest)
		return
	}
	if creds.Password == "" {
		http.Error(w, "Falta el campo contraseña", http.StatusBadRequest)
		return
	}

	// Busca el correo electrónico en la "base de datos" para validar el usuario.
	for _, u := range usersDB {
		if u.Email == creds.Email {
			// Genera un token JWT para el usuario si el correo electrónico existe.
			token, err := GenerateJWT(u)
			if err != nil {
				http.Error(w, "Failed to generate token", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"token": token})
			return
		}
	}

	// Si el correo electrónico no se encuentra en usersDB, devuelve un error.
	http.Error(w, "Email not found", http.StatusUnauthorized)
}
