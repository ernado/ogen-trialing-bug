// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"fmt"

	"github.com/google/uuid"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/Cluster
type Cluster struct {
	ID uuid.UUID `json:"id"`
}

// GetID returns the value of ID.
func (s *Cluster) GetID() uuid.UUID {
	return s.ID
}

// SetID sets the value of ID.
func (s *Cluster) SetID(val uuid.UUID) {
	s.ID = val
}

// DeleteClusterOK is response for DeleteCluster operation.
type DeleteClusterOK struct{}

// Error occurred while processing request.
// Ref: #/components/schemas/Error
type Error struct {
	// Human-readable error message.
	ErrorMessage string `json:"error_message"`
}

// GetErrorMessage returns the value of ErrorMessage.
func (s *Error) GetErrorMessage() string {
	return s.ErrorMessage
}

// SetErrorMessage sets the value of ErrorMessage.
func (s *Error) SetErrorMessage(val string) {
	s.ErrorMessage = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// UpdateClusterOK is response for UpdateCluster operation.
type UpdateClusterOK struct{}
