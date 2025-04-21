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

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/go_course_users")
	if err != nil {
		return nil, err
	}
	return db, nil

}
