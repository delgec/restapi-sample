package model

import (
	"fmt"
	"net/http"
	"restapi/spec"
)

type ErrorSource int

const (
	ErrorNotFound             ErrorSource = http.StatusNotFound
	ErrorSourceServer         ErrorSource = http.StatusInternalServerError
	ErrorSourceClient         ErrorSource = http.StatusBadRequest
	ErrorSourceClientSecurity ErrorSource = http.StatusForbidden
)

type APIError struct {
	Source  ErrorSource
	Code    spec.ErrorCode
	Message string
	Cause   error
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}

	return fmt.Sprintf("%s - %s", e.Code, e.Message)
}
