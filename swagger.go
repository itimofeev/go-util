package util

import (
	"encoding/json"
	"github.com/go-openapi/loads"
	"net/http"
	"net/http/httptest"
)

func NewTransport(h http.Handler) http.RoundTripper {
	return &handlerTransport{h: h}
}

type handlerTransport struct {
	h http.Handler
}

func (s *handlerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp := httptest.NewRecorder()
	s.h.ServeHTTP(resp, req)
	return resp.Result(), nil
}

func ValidateSpec(orig, flat json.RawMessage) *loads.Document {
	swaggerSpec, err := loads.Embedded(orig, flat)
	if err != nil {
		Log.Fatalln(err)
	}
	return swaggerSpec
}
