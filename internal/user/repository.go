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
		Update(ctx context.Context, id string, firstName, lastName, email *string) error // Nuevo método
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
	sqlQ := "INSERT INTO users (first_name,last_name, email) VALUES (?, ?, ?)"
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
	sqlQ := "SELECT id, first_name, last_name, email FROM users WHERE id = ?"
	var u domain.User
	err := r.db.QueryRow(sqlQ, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			r.log.Println("Error al buscar el usuario:", err)
			return nil, response.NotFound("Usuario no encontrado")
		}
		r.log.Println("Error al buscar el usuario:", err)
		return nil, response.InternalServerError("Error al buscar el usuario")
	}

	r.log.Println("Usuario encontrado:", u)

	return &u, nil

}

func (r *repo) Update(ctx context.Context, id string, firstName, lastName, email *string) error {
	var fields []string
	var values []interface{}

	if firstName != nil {
		fields = append(fields, "first_name = ?")
		values = append(values, *firstName)
	}
	if lastName != nil {
		fields = append(fields, "last_name = ?")
		values = append(values, *lastName)
	}
	if email != nil {
		fields = append(fields, "email = ?")
		values = append(values, *email)
	}

	if len(fields) == 0 {
		return response.BadRequest("No se proporcionaron campos para actualizar")
	}

	sqlQ := "UPDATE users SET " + fields[0]
	for i := 1; i < len(fields); i++ {
		sqlQ += ", " + fields[i]
	}
	sqlQ += " WHERE id = ?"
	values = append(values, id)

	res, err := r.db.Exec(sqlQ, values...)
	if err != nil {
		r.log.Println("Error al actualizar el usuario:", err)
		return response.InternalServerError("Error al actualizar el usuario")
	}

	row, errorr := res.RowsAffected()
	if errorr != nil {
		r.log.Println("Error al obtener el número de filas afectadas:", errorr)
		return response.InternalServerError("Error al obtener el número de filas afectadas")
	}

	if row == 0 {
		r.log.Println("No se encontraron filas para actualizar")
		return response.NotFound("No se encontraron filas para actualizar")
	}

	return nil
}
