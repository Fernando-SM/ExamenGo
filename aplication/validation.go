package aplication

import (
	"ExamenGo/auth"
	"ExamenGo/models"
	"errors"
	"regexp"
)

// ValidateUserRegistration verifica si un usuario puede ser registrado.
func ValidateUserRegistration(user models.User) error {
	// Valida el formato del email utilizando ValidateEmail.
	if !ValidateEmail(user.Email) {
		return errors.New("invalid email format")
	}

	// Valida el teléfono asegurándose de que tenga 10 dígitos.
	if !ValidatePhone(user.Phone) {
		return errors.New("phone must be 10 digits long")
	}

	/*	// Valida la contraseña con criterios específicos de longitud, mayúsculas, minúsculas, números y caracteres especiales.
		if !ValidatePassword(user.Password) {
			return errors.New("password does not meet the criteria")
		}
	*/
	// Verifica la unicidad del email y el teléfono en la "base de datos" simulada (UsersDB).
	for _, u := range models.UsersDB {
		if u.Email == user.Email {
			return errors.New("email already registered")
		}
		if u.Phone == user.Phone {
			return errors.New("phone already registered")
		}
	}

	return nil // Devuelve nil si todas las validaciones pasan sin errores.
}

// ValidateUserLogin verifica si las credenciales de usuario son válidas.
func ValidateUserLogin(email, password string) (string, error) {
	// Comprobación específica para un usuario "demo".
	if email == "demo@demo.com" && password == "D3mo123456987_" {
		// Genera un JWT para el usuario si las credenciales son correctas.
		token, err := auth.GenerateJWT(models.User{Email: email}) // Se asume la existencia de una función GenerateJWT.
		if err != nil {
			return "", err // Manejo de errores al generar el JWT.
		}
		return token, nil
	}

	// Devuelve un error si las credenciales no coinciden con el usuario demo.
	return "", errors.New("invalid credentials")
}

// ValidateEmail utiliza una expresión regular para validar el formato del correo electrónico.
func ValidateEmail(email string) bool {
	// La expresión regular verifica que el correo tenga un formato típico con @ y un dominio.
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}

// ValidatePhone utiliza una expresión regular para asegurarse de que el teléfono tenga 10 dígitos.
func ValidatePhone(phone string) bool {
	// La expresión regular verifica que el teléfono sea una secuencia de 10 dígitos.
	re := regexp.MustCompile(`^[0-9]{10}$`)
	return re.MatchString(phone)
}

// ValidatePassword verifica que la contraseña cumpla con los criterios especificados.
func ValidatePassword(password string) bool {
	var (
		hasMinLength = len(password) >= 6                                // Longitud mínima de 6 caracteres.
		hasMaxLength = len(password) <= 12                               // Longitud máxima de 12 caracteres.
		hasUpper     = regexp.MustCompile(`[A-Z]`).MatchString(password) // Al menos una letra mayúscula.
		hasLower     = regexp.MustCompile(`[a-z]`).MatchString(password) // Al menos una letra minúscula.
		hasNumber    = regexp.MustCompile(`[0-9]`).MatchString(password) // Al menos un dígito.
		/*		hasSpecial   = regexp.MustCompile(`[@]`).MatchString(password)   // Al menos un carácter especial (@).
		 */)

	// La contraseña debe cumplir todos los criterios para ser válida.
	return hasMinLength && hasMaxLength && hasUpper && hasLower && hasNumber /*&& hasSpecial*/
}
