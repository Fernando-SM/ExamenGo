package models

import (
	"testing"
)

// TestAddUser verifica que un usuario nuevo se puede agregar correctamente a la base de datos simulada.
func TestAddUser(t *testing.T) {
	// Nuevo usuario para agregar a la base de datos simulada.
	newUser := User{
		Username: "newuser",
		Email:    "newuser@example.com",
		Phone:    "1234567890",
		Password: "newpassword123",
	}

	// Agregar el nuevo usuario a la base de datos simulada.
	UsersDB = append(UsersDB, newUser)

	// Verificar que el usuario esté en la base de datos simulada.
	if UsersDB[len(UsersDB)-1].Email != newUser.Email {
		t.Errorf("New user was not added to the database")
	}
}

// TestFindUser verifica que se puede encontrar un usuario existente por su correo electrónico.
func TestFindUser(t *testing.T) {
	// Correo electrónico del usuario existente en la base de datos simulada.
	emailToFind := "demo@demo.com"

	// Intentar encontrar el usuario en la base de datos simulada.
	var foundUser *User
	for _, u := range UsersDB {
		if u.Email == emailToFind {
			foundUser = &u
			break
		}
	}

	// Verificar que el usuario se encontró correctamente.
	if foundUser == nil || foundUser.Email != emailToFind {
		t.Errorf("User with email %v was not found in the database", emailToFind)
	}
}
