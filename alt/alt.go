package alt

import "net/http"

type DTO[A any] interface {
	Encode(w http.ResponseWriter)
	Decode(r *http.Request) error
}

type HTTPError interface {
	error
	Encode(w http.ResponseWriter)
}

type Service[A DTO[A], B DTO[B]] func(A) (B, HTTPError)

type Handler[A DTO[A], B DTO[B]] struct {
	service Service[A, B]
}

func (h Handler[A, B]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		incomingRequest  A
		outgoingResponse B
		err              HTTPError
	)

	if err := incomingRequest.Decode(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	outgoingResponse, err = h.service(incomingRequest)

	if err != nil {
		err.Encode(w)
		return
	}

	outgoingResponse.Encode(w)

}

