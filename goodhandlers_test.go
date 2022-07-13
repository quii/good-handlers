package goodhandlers_test

import (
	goodhandlers "github.com/quii/good-handlers"
	"github.com/quii/good-handlers/assert"
	"github.com/quii/good-handlers/kyc"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGoodHandlers(t *testing.T) {
	handler := goodhandlers.New(
		kyc.CreateKYC,
		kyc.DecodeRequest,
		kyc.EncodeResponse,
	)
	server := httptest.NewServer(handler)
	defer server.Close()

	postBody := `{"Name": "Chris James", "Age": 38}`

	response, err := http.Post(server.URL, "/", strings.NewReader(postBody))
	assert.NoError(t, err)

	defer response.Body.Close()
	kycResponse, err := kyc.DecodeCreateKYCResponse(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, kycResponse.ID, "lmao")
}
