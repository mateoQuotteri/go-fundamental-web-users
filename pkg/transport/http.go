package transport

import (
	"context"
	"net/http"
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
	//El decode se encargara de decodificar la peticion del cliente y pasarsela al endpoint
	data, err := decode(t.ctx, t.r)

	if err != nil {

		encodeError(t.ctx, t.w, err)
		return
	}

	res, err := endpoint(t.ctx, data)
	if err != nil {

		encodeError(t.ctx, t.w, err)
		return
	}

	encode(t.ctx, t.w, res)
}
