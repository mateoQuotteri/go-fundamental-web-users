package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/mateoQuotteri/go-fundamental-web-users/internal/domain"
	"github.com/mateoQuotteri/go-responses/response"
)

type DB struct {
	Users     []domain.User
	MaxUserID uint64
}

type (
	Repository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
		Get(ctx context.Context, id string) (*domain.User, error)
		Update(ctx context.Context, id string, user *domain.User) error // Nuevo método
	}
	repo struct {
		db  *sql.DB
		log *log.Logger
	}
)

func NewRepository(db *sql.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (r *repo) Create(ctx context.Context, user *domain.User) error {
	sqlQ := "INSERT INTO users (first_name,last_name, email) VALUES (?, ?, ?, ?)"
	res, err := r.db.Exec(sqlQ, user.FirstName, user.LastName, user.Email)

	if err != nil {
		r.log.Println("Error al insertar el usuario:", err)
		return response.InternalServerError("Error al insertar el usuario")
	}

	id, error := res.LastInsertId()
	if error != nil {
		r.log.Println("Error al insertar el usuario:", err)
		return response.InternalServerError("Error al insertar el usuario")
	}
	user.ID = string(id)
	r.log.Println("Usuario insertado:", user, " id: "+user.ID)
	return nil
}

func (r *repo) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	sqlQ := "SELECT id, first_name, last_name, email FROM users"
	rows, err := r.db.Query(sqlQ)
	if err != nil {
		r.log.Println("Error al insertar el usuario:", err)
		return nil, response.InternalServerError("Error al buscar usuarios")
	}

	defer rows.Close()

	//Recorremos los rows
	for rows.Next() {
		var u domain.User
		// Escaneamos los valores de cada fila en la estructura de usuario
		// y los asignamos a la variable u
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)

		if err != nil {
			r.log.Println("Error al escanear el usuario:", err)
			return nil, response.InternalServerError("Error al escanear el usuario en el rows.next")
		}
		users = append(users, u)
	}

	r.log.Println("Usuarios encontrados:", len(users))

	return users, nil
}

func (r *repo) Get(ctx context.Context, id string) (*domain.User, error) {
	/*index := slices.IndexFunc(r.db.Users, func(v domain.User) bool {
		return v.ID == id
	})
	if index < 0 {
		return nil, response.NotFound("Usuario no encontrado, lo lamentamos")
	}
	return &r.db.Users[index], nil*/

	return nil, response.NotFound("Usuario no encontrado, lo lamentamos")
}

func (r *repo) Update(ctx context.Context, id string, user *domain.User) error {
	// Encontrar el índice del usuario que queremos actualizar
	/*index := slices.IndexFunc(r.db.Users, func(v domain.User) bool {
		return v.ID == id
	})

	if index < 0 {
		return &ErrorNotFound{ID: id}
	}

	// Mantener el mismo ID
	user.ID = id

	// Actualizar el usuario en la base de datos
	r.db.Users[index] = *user

	r.log.Println("Usuario actualizado:", user)*/
	return nil
}
