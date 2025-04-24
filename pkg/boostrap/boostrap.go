package boostrap

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "users-api", log.LstdFlags|log.Lshortfile)
}

func NewDB() (*sql.DB, error) {

	dbUrll := os.ExpandEnv("$DATABASE_USER:$DATABASE_PASSWORD@tcp($DATABASE_HOST:$DATABASE_PORT)/$DATABASE_NAME")

	log.Println("Conectando a la base de datos en:", dbUrll)
	db, err := sql.Open("mysql", dbUrll)
	if err != nil {
		return nil, err
	}
	return db, nil

}
