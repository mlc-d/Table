package errs

import "errors"

var (
	ErrBadRequest         = errors.New("invalid request")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
