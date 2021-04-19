package dto

// ErrorResponse is the JSON DTO definition for Error response
type ErrorResponse struct {
	Errors []*Error    `json:"errors"`
	Meta   interface{} `json:"meta"`
}

// NewErrorResponse returns new error response
func NewErrorResponse(errs []*Error) *ErrorResponse {
	return &ErrorResponse{
		Errors: errs,
	}
}

// Error is the standard error DTO definition
type Error struct {
	Message string `json:"message"` // Human-readable error message. Should contain indications about how to fix the error if possible.
}

// NewError returns error DTO
func NewError(message string) *Error {
	return &Error{
		Message: message,
	}
}
