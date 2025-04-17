package transport

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Trasport interface {
	Server(
		endpoint Endpoint,
		//El decode se encargara de decodificar la peticion del cliente y pasarsela al endpoint
		decode func(ctx context.Context, r *http.Request) (interface{}, error),
		//El encode se encargara de responder al cliente con el resultado de la peticion
		encode func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
		//El encodeError se encargara de responder al cliente con un error en caso de que falle la peticion
		encodeError func(ctx context.Context, w http.ResponseWriter, err error),
	)
}

type Endpoint func(ctx context.Context, request interface{}) (interface{}, error)

type trasport struct {
	w   http.ResponseWriter
	r   *http.Request
	ctx context.Context
}

func New(w http.ResponseWriter, r *http.Request, ctx context.Context) Trasport {
	return &trasport{
		w:   w,
		r:   r,
		ctx: ctx,
	}
}

func (t *trasport) Server(
	endpoint Endpoint,
	decode func(ctx context.Context, r *http.Request) (interface{}, error),
	encode func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	encodeError func(ctx context.Context, w http.ResponseWriter, err error),
) {
	log.Println("DEPURACIÓN - Iniciando proceso de Server en transport")

	data, err := decode(t.ctx, t.r)
	if err != nil {
		log.Printf("DEPURACIÓN - Error en decode: %v", err)
		encodeError(t.ctx, t.w, err)
		return
	}

	log.Printf("DEPURACIÓN - Data decodificada: %+v", data)

	// Verificar si endpoint es nil
	if endpoint == nil {
		log.Println("DEPURACIÓN - Endpoint es nil!")
		encodeError(t.ctx, t.w, fmt.Errorf("endpoint is nil"))
		return
	}

	res, err := endpoint(t.ctx, data)
	if err != nil {
		log.Printf("DEPURACIÓN - Error en endpoint: %v", err)
		encodeError(t.ctx, t.w, err)
		return
	}

	log.Printf("DEPURACIÓN - Respuesta del endpoint: %+v", res)

	err = encode(t.ctx, t.w, res)
	if err != nil {
		log.Printf("DEPURACIÓN - Error en encode: %v", err)
	}
}

func Clean(url string) ([]string, int) {
	if url[0] == '/' {
		url = "/" + url
	}

	if url[len(url)-1] == '/' {
		url = url + "/"
	}

	// Recibe un string y lo convierte en slices
	parts := strings.Split(url, "/")

	return parts, len(parts)

}
