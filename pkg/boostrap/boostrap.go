package boostrap

import (
	"log"
	"os"

	"github.com/mateoQuotteri/go-fundamental-web-users/internal/domain"
	"github.com/mateoQuotteri/go-fundamental-web-users/internal/user"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "users-api", log.LstdFlags|log.Lshortfile)
}

func NewDB() user.DB {
	return user.DB{
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
}
