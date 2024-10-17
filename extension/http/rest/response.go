package rest

import (
	"net/http"
)

type Response struct {
	Status int
	header http.Header
	Error  error
	Body   any
}
type ErrorResponse struct {
	Message string `json:"message"`
}
