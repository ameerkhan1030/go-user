package errors

import (
	"fmt"
	"net/http"
	"test/internal/app"

	"github.com/labstack/echo"
)

type handler struct {
	app app.App
}

// RegisterErrorHandler takes in an Echo router and registers routes onto it.
func RegisterErrorHandler(e *echo.Echo, app app.App) {
	h := handler{app}

	e.HTTPErrorHandler = h.handleError
}

// handleError is an Echo error handler that uses HTTP errors accordingly, and any
// generic error will be interpreted as an internal server error.
func (h *handler) handleError(err error, c echo.Context) {
	//log := logger.FromContext(c)

	code := http.StatusInternalServerError
	msg := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = http.StatusText(code)
	}

	// log.WithFields(logrus.Fields{
	// 	"status_code": code,
	// }).WithError(err).Error("request error")

	err = c.JSON(code, map[string]interface{}{"error": map[string]interface{}{"message": msg, "status_code": code}})
	if err != nil {
		fmt.Println("Json Error")
	}
}
