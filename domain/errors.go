package domain

import "net/http"

var (
	// ErrInvalidParam is a RequestError for Bad Request Errors
	ErrInvalidParam = &RequestError{
		ErrorString: "Bad Request",
		ErrorCode:   http.StatusBadRequest,
	}
	// ErrInternalError is a RequestError Internal Errors
	ErrInternalError = &RequestError{
		ErrorString: "Internal Error",
		ErrorCode:   http.StatusInternalServerError,
	}
	// ErrNotFound is a RequestError for Not Found Errors
	ErrNotFound = &RequestError{
		ErrorString: "Request Not Found",
		ErrorCode:   http.StatusNotFound,
		ErrorDetail: "The requested endpoint doesn't exists",
	}

	// ErrLimitParam handle errors when limit parameter is wrong
	ErrLimitParam = &RequestError{
		ErrorString: "Bad Request",
		ErrorCode:   http.StatusBadRequest,
		ErrorDetail: "Parameter 'limit' must be a number and have a value greater than 1",
	}

	// ErrPersonsOrderByParam error
	ErrPersonsOrderByParam = &RequestError{
		ErrorString: "Bad Request",
		ErrorCode:   http.StatusBadRequest,
		ErrorDetail: "Parameter 'orderBy' must be either '-numberOfEpisodes' or 'numberOfEpisodes'",
	}

	// ErrDatabaseConnection error
	ErrDatabaseConnection = &RequestError{
		ErrorString: "Internal Server Error",
		ErrorCode:   http.StatusInternalServerError,
		ErrorDetail: "Problem connecting to database.. Try again later",
	}

	// ErrDatabaseQuery error
	ErrDatabaseQuery = &RequestError{
		ErrorString: "Internal Server Error",
		ErrorCode:   http.StatusInternalServerError,
		ErrorDetail: "Problem when performing database query",
	}
)

// RequestError holds the error message and http code
type RequestError struct {
	ErrorString string
	ErrorCode   int
	ErrorDetail string
}

// Code returns the http error code
func (err *RequestError) Code() int {
	return err.ErrorCode
}

// Error returns the error message
func (err *RequestError) Error() string {
	return err.ErrorString
}

// Detail returns the error detail
func (err *RequestError) Detail() string {
	return err.ErrorDetail
}

// ErrorMessage returns the code and the message
func ErrorMessage(errorType *RequestError) (code int, message map[string]interface{}) {
	return errorType.Code(), map[string]interface{}{
		"message": errorType.Error(),
		"code":    errorType.Code(),
		"detail":  errorType.Detail(),
	}
}
