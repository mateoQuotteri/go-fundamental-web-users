package user

import (
	"context"
	"log"

	"github.com/mateoQuotteri/go-fundamental-web-users/internal/domain"
)

type DB struct {
	Users     []domain.User
	MaxUserID uint64
}

type (
	Repository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
	}

	repo struct {
		db  DB
		log *log.Logger
	}
)

func NewRepository(db DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}

}

func (r *repo) Create(ctx context.Context, user *domain.User) error {
	//Agarro el ultimo id del ultimo usuario en bd y lo incremento en 1 para el nuevo usuario
	r.db.MaxUserID++

	//Seteo el id del nuevo usuario con el id maximo de la bd
	user.ID = string(r.db.MaxUserID)

	//Obtengo la entidad user y le hacemos un append del nuevo user
	r.db.Users = append(r.db.Users, *user)
	r.log.Println("User created:", user)
	return nil
}

func (r *repo) GetAll(ctx context.Context) ([]domain.User, error) {
	//Devuelvo todos los usuarios de la bd
	return r.db.Users, nil
}
