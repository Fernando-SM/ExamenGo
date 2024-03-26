package models

// User representa el modelo de un usuario en el sistema.
type User struct {
	Username string `json:"username"` // Nombre de usuario, utilizado para identificar al usuario de forma única.
	Email    string `json:"email"`    // Dirección de correo electrónico del usuario, utilizada para la comunicación y como identificador alternativo.
	Phone    string `json:"phone"`    // Número de teléfono del usuario, podría usarse para la verificación de la cuenta o comunicaciones.
	Password string `json:"password"` // Contraseña del usuario, utilizada para autenticar al usuario durante el inicio de sesión.
}

// UsersDB simula una base de datos de usuarios registrados en el sistema.
// En un entorno de producción, preferirías utilizar una base de datos real para almacenar esta información,
// pero para propósitos de demostración o desarrollo, una "base de datos" en memoria como esta puede ser útil.
var UsersDB = []User{
	{
		// Un usuario predefinido para demostración o pruebas.
		// Este usuario puede ser utilizado para realizar pruebas de inicio de sesión sin necesidad de registrar un nuevo usuario cada vez.
		Username: "admin",
		Email:    "demo@demo.com",
		Password: "D3mo123456987_",
	},
}
