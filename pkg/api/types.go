package api

import (
	"context"
	"io"
	"net/http"
)

// Wrapper around http.Request with esapi fields
type GenericRequest struct {
	Request *http.Request
	//Path    strings.Builder

	Body io.Reader

	Name       []string
	Format     string
	H          []string
	Help       *bool
	Local      *bool
	S          []string
	V          *bool
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}
