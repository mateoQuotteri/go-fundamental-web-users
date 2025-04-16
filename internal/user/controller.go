package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// Importa el paquete donde está definida la interfaz Service
	// Por ejemplo: "github.com/mateoQuotteri/go-fundamental-web-users/internal/user"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)
	// Elimina la declaración de Service ya que está definida en service.go
	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func MakeEndpoints(ctx context.Context, s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAllUsers(ctx, s, w)
		case http.MethodPost:
			decode := json.NewDecoder(r.Body)
			var req CreateReq
			err := decode.Decode(&req)
			if err != nil {
				MsgResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			PostUser(ctx, s, w, req)
		default:
			InvalidMethod(w)
		}
	}
}

// Ejemplo de función principal
func main() {
	// Aquí deberías inicializar el servicio
	// Por ejemplo: userService := NewUserService()
	ctx := context.Background()

	// Asumiendo que ya tienes un servicio implementado
	var s Service // Deberías inicializar esto con tu implementación real

	http.HandleFunc("/users", MakeEndpoints(ctx, s))

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Esta función es para obtener todos los usuarios
func GetAllUsers(ctx context.Context, s Service, w http.ResponseWriter) {
	//Buscamos usuarios en la capa de servicio
	users, err := s.GetAll(ctx)
	//Si no hay usuarios, devolvemos un mensaje de error
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	DataResponse(w, http.StatusOK, users)
}

func PostUser(ctx context.Context, s Service, w http.ResponseWriter, req CreateReq) {
	if req.FirstName == "" {
		MsgResponse(w, http.StatusBadRequest, "First name is required")
		return
	}
	if req.LastName == "" {
		MsgResponse(w, http.StatusBadRequest, "LastName is required")
		return
	}
	if req.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "Email is required")
		return
	}
	user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	DataResponse(w, http.StatusCreated, user)
}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusMethodNotAllowed
	w.WriteHeader(status)
	fmt.Fprintf(w, "Method Not Allowed")
}

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "msg": "%s"}`, status, message)
}

func DataResponse(w http.ResponseWriter, status int, data interface{}) {
	//Marshal me trata de convertir una entidad a un json
	value, err := json.Marshal(data)
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, string(value))
}
