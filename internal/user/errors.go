package user

import (
	"errors"
	"fmt"
)

var ErrFirstNameRequired = errors.New("El nombre es requerido")
var ErrLastNameRequired = errors.New("El apellido es requerido")
var ErrEmailRequired = errors.New("El email es requerido")

type ErrorNotFound struct {
	ID string
}

func (e *ErrorNotFound) Error() string {
	return fmt.Sprintf("Usuario con ID %s no encontrado", e.ID)
}
