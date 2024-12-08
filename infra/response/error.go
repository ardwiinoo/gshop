package response

import (
	"errors"
	"net/http"
)

// error general
var (
	ErrNotFound = errors.New("not found")
)

var (
	ErrEmailRequired = errors.New("email is required")
	ErrEmailInvalid = errors.New("email is invalid")
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must be at least 8 characters")
	ErrAuthIsNotExists = errors.New("auth is not exists")
	ErrAuthIsExists = errors.New("auth is exists")
	ErrEmailAlreadyUsed = errors.New("email already used")
	ErrPasswordNotMatch = errors.New("password not match")
	ErrProductRequired = errors.New("product is required")
	ErrProductNameInvalid = errors.New("product name must be at least 3 characters")
	ErrStockInvalid = errors.New("stock must be greater than 0")
	ErrPriceInvalid = errors.New("price must be greater than 0")
	ErrAmountInvalid = errors.New("invalid amount")
	ErrAmountGreaterThanStock = errors.New("amount greater than stock")
)

type Error struct {
	Message string
	Code    string
	HttpCode int
}

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message: msg,
		Code:    code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral = NewError("general error", "99999", http.StatusInternalServerError) 
	ErrorBadRequest = NewError("bad request", "40000", http.StatusBadRequest)
	ErrorNotFound = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
)

var (
	ErrorEmailRequired = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)
	ErrorProductRequired = NewError(ErrProductRequired.Error(), "40005", http.StatusBadRequest)
	ErrorProductNameInvalid = NewError(ErrProductNameInvalid.Error(), "40006", http.StatusBadRequest)
	ErrorStockInvalid = NewError(ErrStockInvalid.Error(), "40007", http.StatusBadRequest)
	ErrorPriceInvalid = NewError(ErrPriceInvalid.Error(), "40008", http.StatusBadRequest)
	ErrorInvalidAmount = NewError(ErrAmountInvalid.Error(), "40009", http.StatusBadRequest)

	ErrorAuthIsNotExists = NewError(ErrAuthIsNotExists.Error(), "40401", http.StatusNotFound)
	ErrorEmailAlreadyUsed = NewError(ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)
	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)
)

var (
	ErrorMapping = map[string]Error{
		ErrNotFound.Error(): ErrorNotFound,
		ErrEmailRequired.Error(): ErrorEmailRequired,
		ErrEmailInvalid.Error(): ErrorEmailInvalid,
		ErrPasswordRequired.Error(): ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
		ErrAuthIsNotExists.Error(): ErrorAuthIsNotExists,
		ErrEmailAlreadyUsed.Error(): ErrorEmailAlreadyUsed,
		ErrPasswordNotMatch.Error(): ErrorPasswordNotMatch,
	}
)