package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mateoQuotteri/go-fundamental-web-users/internal/user"
	"github.com/mateoQuotteri/go-fundamental-web-users/pkg/transport"
)

func NewUserHTTPServer(ctx context.Context, router *http.ServeMux, endpoints user.Endpoints) {
	router.HandleFunc("/users", UserServer(ctx, endpoints))
}

func UserServer(ctx context.Context, endpoints user.Endpoints) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tran := transport.New(w, r, ctx)
		switch r.Method {
		case http.MethodGet:
			tran.Server(
				transport.Endpoint(endpoints.GetAll),
				decodeGetAllUser,
				encodeResponse,
				encodeError)
			return
		default:
			InvalidMethod(w)
		}
	}
}

func decodeGetAllUser(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// Configurar headers antes de escribir el status
	w.Header().Set("Content-Type", "application/json")

	// Marshal convierte una entidad a JSON
	data, err := json.Marshal(response) // Cambiado 'data' por 'response'
	if err != nil {
		return err
	}

	status := http.StatusOK
	w.WriteHeader(status)

	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, string(data))
	return nil
}

func encodeError(ctx context.Context, w http.ResponseWriter, err error) {
	// Configurar headers antes de escribir el status
	w.Header().Set("Content-Type", "application/json")

	status := http.StatusInternalServerError
	w.WriteHeader(status)

	fmt.Fprintf(w, `{"status": %d, "error": "%s"}`, status, err.Error())
}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusMethodNotAllowed
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Method Not Allowed")
}
