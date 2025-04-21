package user

import (
	"context"
	"log"

	"github.com/mateoQuotteri/go-fundamental-web-users/internal/domain"
)

type (
	Service interface {
		Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error)
		GetAll(ctx context.Context) ([]domain.User, error)
		Get(ctx context.Context, id string) (*domain.User, error)
		Update(ctx context.Context, id, firstName, lastName, email string) (*domain.User, error) // Nuevo método
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(repo Repository, l *log.Logger) Service {
	return &service{
		repo: repo,
		log:  l,
	}
}

func (s *service) Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error) {
	//Creo el nuevo usuario
	user := &domain.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
	//Guardo el nuevo usuario en la base de datos
	// pasnadole el conontexto que recibimos desde el servicio y el usuario que generamos
	err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	s.log.Println("Usuario creadoo:", user)
	return user, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.User, error) {
	//Obtengo todos los usuarios de la base de datos
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	s.log.Println("Usuarios obtenidos:", users)
	return users, nil
}

func (s *service) Get(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) Update(ctx context.Context, id, firstName, lastName, email string) (*domain.User, error) {
	// Verificar si el usuario existe
	_, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Crear la estructura de usuario actualizada
	updatedUser := &domain.User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	// Actualizar el usuario en el repositorio
	err = s.repo.Update(ctx, id, updatedUser)
	if err != nil {
		return nil, err
	}

	s.log.Println("Usuario actualizado:", updatedUser)
	return updatedUser, nil
}
