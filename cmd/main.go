package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mateoQuotteri/go-fundamental-web-users/internal/user"
	"github.com/mateoQuotteri/go-fundamental-web-users/pkg/boostrap"
)

func main() {
	_ = godotenv.Load()

	db, err := boostrap.NewDB()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	errore := db.Ping()
	if errore != nil {
		log.Fatal("Error al hacer ping a la base de datos:", errore)
	}

	logger := boostrap.NewLogger()
	repo := user.NewRepository(db, logger)
	_ = user.NewService(repo, logger)
	_ = context.Background()
	// Inicializar el servidor HTTP
	// Define the accessControl function
	accessControl := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Add your access control logic here
			h.ServeHTTP(w, r)
		})
	}

	// Crear un router (handler)
	router := http.NewServeMux()

	// Aquí deberías configurar tus rutas
	// Por ejemplo:
	// router.HandleFunc("/users", userHandler)

	// Definir el puerto (puedes obtenerlo de una variable de entorno)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080" // Puerto por defecto si no está definido
	} else {
		port = ":" + port // Añadir : al inicio del puerto
	}

	addr := fmt.Sprintf("127.0.0.1%s", port)
	fmt.Println("Servidor iniciando en", addr)

	// Configurar el servidor HTTP
	server := &http.Server{
		Handler: accessControl(router),
		Addr:    addr,
	}

	// Iniciar el servidor
	log.Fatal(server.ListenAndServe())
}
