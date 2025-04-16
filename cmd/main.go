package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mateoQuotteri/go-fundamental-web-users/internal/domain"
	"github.com/mateoQuotteri/go-fundamental-web-users/internal/user"
)

func main() {
	server := http.NewServeMux()

	db := user.DB{
		Users: []domain.User{{
			ID:        "1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "mateo@gmail.com",
		},
			{
				ID:        "2",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "mateo@gmail.com",
			},
			{
				ID:        "3",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "mateo@gmail.com",
			}},

		MaxUserID: 3,
	}

	logger := log.New(os.Stdout, "users-api", log.LstdFlags|log.Lshortfile)
	repo := user.NewRepository(db, logger)
	service := user.NewService(repo, logger)
	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoints(ctx, service))
	fmt.Println("Hello, World!")

	log.Fatal(http.ListenAndServe(":8080", server))
}
