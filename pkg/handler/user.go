package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateoQuotteri/go-fundamental-web-users/internal/user"
	"github.com/mateoQuotteri/go-fundamental-web-users/pkg/transport"
	"github.com/mateoQuotteri/go-responses/response"
)

// Esta funcion es la encargada de inicializar el servidor HTTP para el servicio de usuario
// Recibe los endpoints del servicio de usuario
// y configura el router para manejar las solicitudes HTTP
// para el servicio de usuario
func NewUserHTTPServer(endpoints user.Endpoints) *gin.Engine {
	r := gin.Default()

	r.POST("/users", transport.GinServer(
		endpoints.Create,
		decodeCreateUser,
		encodeResponse,
		encodeError,
	))

	r.GET("/users", transport.GinServer(
		endpoints.GetAll,
		decodeGetAllUser,
		encodeResponse,
		encodeError,
	))

	r.GET("/users/:userID", transport.GinServer(
		endpoints.Get,
		decodeGetUser,
		encodeResponse,
		encodeError,
	))

	r.PUT("/users/:userID", transport.GinServer(
		endpoints.Update,
		decodeUpdateUser,
		encodeResponse,
		encodeError,
	))

	return r
}

func decodeGetAllUser(c *gin.Context) (interface{}, error) {
	return nil, nil
}

func decodeGetUser(c *gin.Context) (interface{}, error) {
	userID := c.Param("userID")
	if userID == "" {
		return nil, fmt.Errorf("userID not found in parameters")
	}

	return user.GetReq{
		ID: userID,
	}, nil
}

func decodeCreateUser(c *gin.Context) (interface{}, error) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		return nil, response.Unauthorized("Unauthorized")
	}

	var req user.CreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, fmt.Errorf("invalid request format: %v", err.Error())
	}

	return req, nil
}

// Adaptado para Gin
func decodeUpdateUser(c *gin.Context) (interface{}, error) {
	userID := c.Param("userID")
	if userID == "" {
		return nil, fmt.Errorf("userID not found in parameters")
	}

	var req user.UpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, fmt.Errorf("invalid request format: %v", err.Error())
	}

	// Asegúrate de que el ID del path coincida con el de la solicitud
	req.ID = userID
	return req, nil
}

// Actualizado para Gin
func encodeResponse(c *gin.Context, response interface{}) error {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   response,
	})

	return nil
}

// Actualizado para Gin
func encodeError(c *gin.Context, err error) error {
	status := http.StatusInternalServerError

	// Si el error es de tipo response, podríamos manejar códigos personalizados
	// Por ejemplo, verificar si es un error de tipo response.Unauthorized

	c.JSON(status, gin.H{
		"status": status,
		"error":  err.Error(),
	})

	return nil
}
