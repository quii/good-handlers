package goodhandlers

import (
	"context"
	"io"
	"net/http"
)

type Decoder[T any] func(io.Reader) (T, error)
type Encoder[T any] func(T, io.Writer) error

// this is what every "service" layer should look like, defined as types
type Service[A, B any] func(context.Context, A) (B, error)

func New[A, B any](service Service[A, B], decode Decoder[A], encode Encoder[B]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := decode(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//todo: error handling might be difficult to do nicely, doing 500 as a default is whack
		//todo: what about response headers?

		response, err := service(r.Context(), body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		encode(response, w)
		w.WriteHeader(http.StatusOK) // holy assumption batman
	})
}
