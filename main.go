package main

import (
	"github.com/gorilla/mux" // Importa el paquete mux de Gorilla para manejar rutas HTTP.
	"log"                    // Importa el paquete log para registrar mensajes.
	"net/http"               // Importa el paquete net/http para utilizar funcionalidades HTTP.
)

func main() {
	router := mux.NewRouter() // Crea una nueva instancia de router de Gorilla Mux.

	// Define las rutas y los manejadores.
	// Ruta de registro: llama a UserRegisterHandler cuando se realiza una solicitud POST a "/register".
	router.HandleFunc("/register", UserRegisterHandler).Methods("POST")

	// Ruta de login: llama a UserLoginHandler cuando se realiza una solicitud POST a "/login".
	router.HandleFunc("/login", UserLoginHandler).Methods("POST")

	// Inicia el servidor.
	// Logra un mensaje indicando que el servidor está iniciando y en qué puerto está escuchando.
	log.Println("Server starting on :8089...")

	// Inicia el servidor en el puerto 8089 y pasa el router como el manejador.
	// log.Fatal se utiliza para registrar un mensaje fatal y terminar el programa si http.ListenAndServe devuelve un error.
	log.Fatal(http.ListenAndServe(":8089", router))
}
