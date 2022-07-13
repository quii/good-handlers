package kyc

import (
	"context"
	"encoding/json"
	"io"
	"time"
)

func CreateKYC(ctx context.Context, kycdto CreateKYCDTO) (CreateKYCResponse, error) {
	// this is where your amazing business logic goes, it doesn't care about HTTP,
	// its just got some lump of data and it needs to return something back

	// in our case, it generates an ID, thrilling
	return CreateKYCResponse{
		ID:        "lmao",
		CreatedAt: time.Now(),
	}, nil
}

type CreateKYCDTO struct {
	Name string
	Age  uint8
}

func DecodeCreateKYCDTO(in io.Reader) (CreateKYCDTO, error) {
	var out CreateKYCDTO
	err := json.NewDecoder(in).Decode(&out)
	return out, err
}

type CreateKYCResponse struct {
	ID        string
	CreatedAt time.Time
}

func DecodeCreateKYCResponse(in io.Reader) (CreateKYCResponse, error) {
	var out CreateKYCResponse
	err := json.NewDecoder(in).Decode(&out)
	return out, err
}

func EncodeOutgoingKYCDTO(res CreateKYCResponse, out io.Writer) error {
	return json.NewEncoder(out).Encode(res)
}
