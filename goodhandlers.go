package goodhandlers

import (
	"io"
	"net/http"
)

type Decoder[T any] func(io.Reader) (T, error)
type Encoder[T any] func(T, io.Writer) error

// this is what every "service" layer should look like, defined as types
type Service[A, B any] func(A) (B, error)

// still returns http.handler, so can use with chi, gorilla or just stdlib
func New[A, B any](service Service[A, B], decode Decoder[A], encode Encoder[B]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := decode(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response, err := service(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		encode(response, w)
	})
}
