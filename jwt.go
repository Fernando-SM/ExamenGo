package main

import (
	"github.com/golang-jwt/jwt" // Importa el paquete jwt para trabajar con JSON Web Tokens.
	"time"
)

var jwtKey = []byte("your_secret_key") // Clave secreta utilizada para firmar los tokens JWT.

// Claims es una estructura personalizada que incluye el nombre de usuario y hereda de jwt.StandardClaims.
type Claims struct {
	Username           string `json:"username"` // Nombre de usuario que se incluirá en el token JWT.
	jwt.StandardClaims        // Campos estándar de un JWT, como tiempos de expiración, emisor, etc.
}

// GenerateJWT genera un nuevo token JWT para un usuario dado.
func GenerateJWT(user User) (string, error) {
	// Define el tiempo de expiración del token a una hora desde el momento actual.
	expirationTime := time.Now().Add(1 * time.Hour)

	// Crea las reclamaciones (claims) del token, incluyendo el nombre de usuario y el tiempo de expiración.
	claims := &Claims{
		Username: user.Username, // Nombre de usuario extraído del objeto User proporcionado.
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // Tiempo de expiración expresado como un timestamp Unix.
		},
	}

	// Crea el token JWT con las reclamaciones especificadas y el método de firma HS256.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firma el token con la clave secreta y lo convierte a string.
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// Devuelve un error si la firma falla.
		return "", err
	}

	// Devuelve el token JWT como string si todo ha ido bien.
	return tokenString, nil
}
