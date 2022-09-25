package model

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
}

func (r *Request) Check() error {
	if r.Method != http.MethodGet &&
		r.Method != http.MethodHead {
		return errors.New("method is not valid")
	}

	_, err := url.ParseRequestURI(r.URL)
	if err != nil {
		return errors.New("URL is not valid")
	}

	return nil
}

type Response struct {
	ID      uuid.UUID
	Status  int
	Headers map[string][]string
	Length  int64
}
