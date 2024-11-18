package response

import "errors"

var (
	ErrEmailRequired = errors.New("email is required")
	ErrEmailInvalid = errors.New("email is invalid")
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordLength = errors.New("password must be at least 8 characters")
)