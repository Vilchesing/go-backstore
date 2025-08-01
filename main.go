package main

import (
	"database/sql" // Necesitas importar sql aquí también para el tipo *sql.DB
	"fmt"
	"log"
	"net/http" // Un ejemplo si estás creando un servidor web
	"vilchesing/go-backstore/config"

	"github.com/joho/godotenv"
	// Importa tu paquete 'config'
	// Reemplaza "tu_modulo_go" con la ruta real de tu módulo.
	// Por ejemplo, si tu archivo go.mod dice "module github.com/tu_usuario/tu_app",
	// entonces la ruta sería "github.com/tu_usuario/tu_app/config".
)

// Declara una variable global para la conexión a la base de datos.
// Es común tener la conexión accesible para el resto de tu aplicación.
var db *sql.DB

func main() {
	// Inicializar la conexión a la base de datos es una de las primeras cosas
	// que debe hacer tu aplicación.
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env. Asegúrate de que las variables de entorno estén configuradas de otra manera.")
	}
	db, err = config.InitDB() // ¡Aquí es donde llamas a tu función InitDB!
	if err != nil {
		// Si hay un error al inicializar la DB, la aplicación no puede continuar.
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	// Muy importante: Asegúrate de cerrar la conexión a la base de datos
	// cuando la función main termine de ejecutarse.
	defer func() {
		if db != nil {
			err := db.Close()
			if err != nil {
				log.Printf("Error al cerrar la conexión de la base de datos: %v", err)
			} else {
				fmt.Println("Conexión a la base de datos cerrada exitosamente.")
			}
		}
	}()

	fmt.Println("¡Aplicación iniciada! La conexión a la base de datos está lista para usar.")

	// --- Ejemplo de uso de la conexión en tu aplicación (por ejemplo, un servidor web) ---
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/usuarios", getUsersHandler)

	fmt.Println("Servidor escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// homeHandler es una función de ejemplo para una ruta básica.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Bienvenido a mi aplicación!")
}

// getUsersHandler es una función de ejemplo que usaría la conexión 'db'.
func getUsersHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Accedido al endpoint de usuarios. La conexión a la DB está disponible para usar.")
}
