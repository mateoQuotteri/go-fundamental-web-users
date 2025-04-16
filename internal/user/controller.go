package user

import (
	"context"
	"fmt"
)

type (
	// Controller es una función que procesa una solicitud y devuelve una respuesta o un error
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	// Endpoints contiene todos los endpoints del servicio de usuario
	Endpoints struct {

		//Create y GetAll son funciones que manejan las solicitudes HTTP
		// y que respetan este patron (linea 10) func(ctx context.Context, request interface{}) (interface{}, error)
		Create Controller
		GetAll Controller
	}

	// CreateReq estructura para la solicitud de creación de usuario
	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

// MakeEndpoints crea todos los endpoints para el servicio de usuario
func MakeEndpoints(ctx context.Context, s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
	}
}

// makeGetAllEndpoint crea un endpoint para obtener todos los usuarios
func makeGetAllEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// Buscamos usuarios en la capa de servicio
		users, err := s.GetAll(ctx)
		// Si hay un error, lo devolvemos
		if err != nil {
			return nil, err
		}
		return users, nil
	}
}

// makeCreateEndpoint crea un endpoint para crear un usuario
func makeCreateEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateReq)
		if !ok {
			return nil, fmt.Errorf("invalid request format")
		}

		// Validaciones
		if req.FirstName == "" {
			return nil, fmt.Errorf("first name is required")
		}
		if req.LastName == "" {
			return nil, fmt.Errorf("last name is required")
		}
		if req.Email == "" {
			return nil, fmt.Errorf("email is required")
		}

		// Crear usuario
		user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}
