package response

import "errors"

// error general
var (
	ErrNotFound = errors.New("not found")
)

var (
	ErrEmailRequired = errors.New("email is required")
	ErrEmailInvalid = errors.New("email is invalid")
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordLength = errors.New("password must be at least 8 characters")
	ErrAuthIsNotExists = errors.New("auth is not exists")
	ErrAuthIsExists = errors.New("auth is exists")
	ErrEmailAlreadyUsed = errors.New("email already used")
)