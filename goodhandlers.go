package goodhandlers

import (
	"context"
	"net/http"
)

type RequestDecoder[A any] func(r *http.Request) (A, error)
type ResponseEncoder[B any] func(http.ResponseWriter, B, error)
type Service[A, B any] func(context.Context, A) (B, error)

func New[A, B any](service Service[A, B], decode RequestDecoder[A], encode ResponseEncoder[B]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		request, err := decode(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := service(r.Context(), request)
		encode(w, response, err)
	})
}
