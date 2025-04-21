package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mateoQuotteri/go-fundamental-web-users/internal/user"
	"github.com/mateoQuotteri/go-fundamental-web-users/pkg/boostrap"
	"github.com/mateoQuotteri/go-fundamental-web-users/pkg/handler"
)

func main() {
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

	handler.NewUserHTTPServer(ctx, server, user.MakeEndpoints(ctx, service))

	fmt.Println("Hello, World!")

	log.Fatal(http.ListenAndServe(":8080", server))
}
