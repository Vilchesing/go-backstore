package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB inicializa y devuelve un pool de conexiones a la base de datos.
// Configura parámetros importantes del pool como el número máximo de conexiones.
func InitDB() (*sql.DB, error) {
	// Construye el DSN (Data Source Name) a partir de variables de entorno

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// --- TEMPORAL: Imprime los valores para depurar ---
	fmt.Println("DEBUG: DB_USER =", os.Getenv("DB_USER"))
	fmt.Println("DEBUG: DB_PASSWORD (debe estar vacío si no hay) =", os.Getenv("DB_PASSWORD"))
	fmt.Println("DEBUG: DB_HOST =", os.Getenv("DB_HOST"))
	fmt.Println("DEBUG: DB_PORT =", os.Getenv("DB_PORT"))
	fmt.Println("DEBUG: DB_NAME =", os.Getenv("DB_NAME"))
	fmt.Println("DEBUG: DSN generado =", dsn)

	// Abre la conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("fallo al abrir la conexión a la base de datos: %w", err)
	}

	// --- Configura el Pool de Conexiones (¡SÚPER IMPORTANTE!) ---
	// Se Define el número máximo de conexiones que pueden estar abiertas a la vez.
	db.SetMaxOpenConns(25) // Ejemplo: 25 conexiones abiertas como máximo

	// Define el número máximo de conexiones inactivas (en espera) en el pool.
	// Debería ser menor o igual a MaxOpenConns.
	db.SetMaxIdleConns(10) // Ejemplo: 10 conexiones inactivas listas para usar

	// Verifica que la conexión funcione haciendo un "ping" a la base de datos
	if err = db.Ping(); err != nil {
		db.Close() // ¡Importante! Cierra la conexión si el ping falla
		return nil, fmt.Errorf("fallo al conectar a la base de datos (ping fallido): %w", err)
	}

	fmt.Println("¡Conexión con la base de datos establecida exitosamente!")
	return db, nil // Devuelve la conexión establecida
}
