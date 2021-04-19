package validation

const validationError = "Validation errors"

// Error is used for API validation errors
type Error struct {
	field   string
	message string
}

// Error returns error message
func (e *Error) Error() string {
	return e.message
}

// Field indicates which field is not validated
func (e *Error) Field() string {
	return e.field
}

// New returns new validation error
func New(field, message string) *Error {
	return &Error{
		field:   field,
		message: message,
	}
}

// ErrorGroup implements error interface and is used to collect several errors
type ErrorGroup struct {
	errors []*Error
}

// NewErrorGroup returns new ErrorGroup
func NewErrorGroup() *ErrorGroup {
	return &ErrorGroup{
		errors: []*Error{},
	}
}

func (e *ErrorGroup) Error() string {
	return validationError
}

// GetErrors returns group error list
func (e *ErrorGroup) GetErrors() []*Error {
	return e.errors
}

// HasErrors returns true if group has non empty error list
func (e *ErrorGroup) HasErrors() bool {
	return len(e.errors) != 0
}

// AddError adds error to error group
func (e *ErrorGroup) AddError(err *Error) {
	e.errors = append(e.errors, err)
}
