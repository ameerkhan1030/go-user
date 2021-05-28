package response

import (
	"net/http"

	"github.com/labstack/echo"
)

// Response represents standardized structure for all API responses.
type Response struct {
	StatusCode int         `json:"statusCode"`
	Payload    interface{} `json:"payload,omitempty"`
	ErrorCode  int         `json:"errorCode,omitempty"`
	Message    string      `json:"message,omitempty"`
}

// OK is used for HTTP status code 200.
func OK(c echo.Context, payload interface{}, message string) error {
	status := http.StatusOK
	return c.JSON(status, Response{
		StatusCode: status,
		Payload:    payload,
		Message:    message,
	})
}

// Unauthorized is used for HTTP status code 401.
func Unauthorized(c echo.Context, payload interface{}, message string, errCode int) error {
	status := http.StatusUnauthorized
	return c.JSON(status, Response{
		StatusCode: status,
		Payload:    payload,
		Message:    message,
		ErrorCode:  errCode,
	})
}

// Forbidden is used for HTTP status code 403.
func Forbidden(c echo.Context, payload interface{}, message string, errCode int) error {
	status := http.StatusForbidden
	return c.JSON(status, Response{
		StatusCode: status,
		Payload:    payload,
		Message:    message,
		ErrorCode:  errCode,
	})
}

// BadRequest is used for HTTP status code 400.
func BadRequest(c echo.Context, payload interface{}, message string, errCode int) error {
	status := http.StatusBadRequest
	return c.JSON(status, Response{
		StatusCode: status,
		Payload:    payload,
		Message:    message,
		ErrorCode:  errCode,
	})
}

// NotFound is used for HTTP status code 404.
func NotFound(c echo.Context, payload interface{}, message string, errCode int) error {
	status := http.StatusNotFound
	return c.JSON(status, Response{
		StatusCode: status,
		Payload:    payload,
		Message:    message,
		ErrorCode:  errCode,
	})
}

// InternalServerError is used for HTTP status code 500.
func InternalServerError(c echo.Context, payload interface{}, message string, errCode int) error {
	status := http.StatusInternalServerError
	return c.JSON(status, Response{
		StatusCode: status,
		Payload:    payload,
		Message:    message,
		ErrorCode:  errCode,
	})
}

// Error is used return with specified message, error code and HTTP status code.
func Error(c echo.Context, status int, payload interface{}, message string, errCode int) error {
	return c.JSON(status, Response{
		StatusCode: status,
		Payload:    payload,
		Message:    message,
		ErrorCode:  errCode,
	})
}