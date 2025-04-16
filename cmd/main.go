package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mateoQuotteri/go-fundamental-web-users/internal/user"
	"github.com/mateoQuotteri/go-fundamental-web-users/pkg/boostrap"
)

func main() {
	server := http.NewServeMux()

	db := boostrap.NewDB()

	logger := boostrap.NewLogger()
	repo := user.NewRepository(db, logger)
	service := user.NewService(repo, logger)
	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoints(ctx, service))
	fmt.Println("Hello, World!")

	log.Fatal(http.ListenAndServe(":8080", server))
}
