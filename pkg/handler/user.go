package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mateoQuotteri/go-fundamental-web-users/internal/user"
	"github.com/mateoQuotteri/go-fundamental-web-users/pkg/transport"
)

// Esta funcion es la encargada de inicializar el servidor HTTP para el servicio de usuario
// Recibe un contexto, un router y los endpoints del servicio de usuario
// y configura el router para manejar las solicitudes HTTP
// para el servicio de usuario
func NewUserHTTPServer(ctx context.Context, router *http.ServeMux, endpoints user.Endpoints) {
	router.HandleFunc("/users/", UserServer(ctx, endpoints))
}

// UserServer es la función que maneja las solicitudes HTTP para el servicio de usuario
// Recibe un contexto y los endpoints del servicio de usuario
// y devuelve una función que maneja las solicitudes HTTP
func UserServer(ctx context.Context, endpoints user.Endpoints) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		log.Printf("Request URL: %s", url)
		log.Printf("Request Method: %s", r.Method)
		path, pathSize := transport.Clean(url)
		log.Printf("Path size: %d", pathSize)
		log.Printf("Path content: %v", path)
		if pathSize < 1 || pathSize > 6 {
			log.Printf("Invalid path size: %d", pathSize)
			InvalidMethod(w)
			return
		}
		params := make(map[string]string)
		if pathSize == 4 && path[2] != "" {
			params["userID"] = path[3]
		}
		ctx := context.WithValue(ctx, "params", params)
		// el transport es el encargado de manejar las solicitudes HTTP
		// y de decodificar y codificar las respuestas
		// y los errores
		tran := transport.New(w, r, ctx)
		switch r.Method {
		case http.MethodGet:
			switch pathSize {
			case 5:
				tran.Server(
					transport.Endpoint(endpoints.GetAll),
					decodeGetAllUser,
					encodeResponse,
					encodeError)
				return
			case 4:
				tran.Server(
					nil,
					decodeGetUser,
					encodeResponse,
					encodeError)
				return
			}
		case http.MethodPost:
			switch pathSize {
			case 5:
				tran.Server(
					transport.Endpoint(endpoints.Create),
					decodeCreateUser,
					encodeResponse,
					encodeError)
				return
			}
		default:
			InvalidMethod(w)
		}
	}
}
func decodeGetAllUser(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
func decodeGetUser(ctx context.Context, r *http.Request) (interface{}, error) {
	params := ctx.Value("params").(map[string]string)
	log.Printf("DEPURACIÓN - Parámetros: %v", params) // Muestra el mapa completo en lugar de solo "params users"
	return nil, fmt.Errorf("myerror")
}
func decodeCreateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req user.CreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid request format: %v", err.Error())
	}
	return req, nil
}

// En handler.go, modifica encodeResponse
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
