package pixelit

// Err implements the error interface so we can have constant errors.
type Err string

func (e Err) Error() string {
	return string(e)
}

// Error Constants
const (
	ErrBadRequest    = Err("pixelit: bad request")
	ErrUnauthorized  = Err("pixelit: unauthorized request")
	ErrForbidden     = Err("pixelit: forbidden")
	ErrNotFound      = Err("pixelit: not found")
	ErrInternalError = Err("pixelit: internal error")
	ErrUnknown       = Err("pixelit: unexpected error occurred")
)
