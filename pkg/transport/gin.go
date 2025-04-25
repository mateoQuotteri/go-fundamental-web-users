package transport

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func GinServer(
	endpoint func(context.Context, interface{}) (interface{}, error),
	decode func(c *gin.Context) (interface{}, error),
	encode func(c *gin.Context, response interface{}) error,
	encodeError func(c *gin.Context, err error) error,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("DEPURACIÓN - Iniciando proceso de Server en transport")

		data, err := decode(c)
		if err != nil {
			log.Printf("DEPURACIÓN - Error en decode: %v", err)
			encodeError(c, err)
			return
		}

		log.Printf("DEPURACIÓN - Data decodificada: %+v", data)

		// Verificar si endpoint es nil
		if endpoint == nil {
			log.Println("DEPURACIÓN - Endpoint es nil!")
			encodeError(c, fmt.Errorf("endpoint is nil"))
			return
		}

		res, err := endpoint(c.Request.Context(), data)
		if err != nil {
			log.Printf("DEPURACIÓN - Error en endpoint: %v", err)
			encodeError(c, err)
			return
		}

		log.Printf("DEPURACIÓN - Respuesta del endpoint: %+v", res)

		err = encode(c, res)
		if err != nil {
			log.Printf("DEPURACIÓN - Error en encode: %v", err)
			encodeError(c, err)
		}
	}
}
