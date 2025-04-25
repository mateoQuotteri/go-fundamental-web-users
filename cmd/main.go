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
	"github.com/mateoQuotteri/go-fundamental-web-users/pkg/handler"
)

func main() {

	_ = godotenv.Load()

	server := http.NewServeMux()

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
	service := user.NewService(repo, logger)
	ctx := context.Background()

	// Inicializar el servidor HTTP

	h := handler.NewUserHTTPServer(user.MakeEndpoints(ctx, service))

	port := os.Getenv("PORT")

	addr := fmt.Sprintf("127.0.0.1%s", port)

	fmt.Println("Hello, World!")
	server := &http.Server{
		Handler: accesControl(h),
		addr:    addr,
	}

	log.Fatal(server.ListenAndServe())
}
