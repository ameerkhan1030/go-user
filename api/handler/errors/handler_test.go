package errors

import (
	"errors"
	"net/http"
	"testing"

	"bitbucket.org/team_carevisor/proto-oob-handler/internal/app"
	"bitbucket.org/team_carevisor/proto-oob-handler/internal/test"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

const errText = "this shouldn't be sent to customers"

func TestHandler(t *testing.T) {
	app := app.App{}
	h := handler{app}

	t.Run("surfaces generic errors as internal server errors", func(tt *testing.T) {
		c, rr := test.NewContext(t, nil)
		err := errors.New("foo")

		h.handleError(err, c)

		assert.Equal(tt, http.StatusInternalServerError, rr.Code)
		assert.Contains(tt, rr.Body.String(), "Internal Server Error")
	})

	t.Run("surfaces HTTP errors transparently but obfuscates message", func(tt *testing.T) {
		c, rr := test.NewContext(t, nil)
		err := echo.NewHTTPError(http.StatusForbidden, "foo")

		h.handleError(err, c)

		assert.Equal(tt, http.StatusForbidden, rr.Code)
		assert.Contains(tt, rr.Body.String(), "Forbidden")
	})

	t.Run("overwrites HTTP 400 error messages", func(tt *testing.T) {
		c, rr := test.NewContext(t, nil)
		err := echo.NewHTTPError(http.StatusBadRequest, errText)

		h.handleError(err, c)

		assert.Equal(tt, http.StatusBadRequest, rr.Code)
		assert.Contains(tt, rr.Body.String(), "Bad Request")
	})

	t.Run("overwrites HTTP 403 error messages", func(tt *testing.T) {
		c, rr := test.NewContext(t, nil)
		err := echo.NewHTTPError(http.StatusForbidden, errText)

		h.handleError(err, c)

		assert.Equal(tt, http.StatusForbidden, rr.Code)
		assert.Contains(tt, rr.Body.String(), "Forbidden")
	})

	t.Run("overwrites HTTP 404 error messages", func(tt *testing.T) {
		c, rr := test.NewContext(t, nil)
		err := echo.NewHTTPError(http.StatusNotFound, errText)

		h.handleError(err, c)

		assert.Equal(tt, http.StatusNotFound, rr.Code)
		assert.Contains(tt, rr.Body.String(), "Not Found")
	})
}

func TestRegisterErrorHandler(t *testing.T) {
	app := app.App{}

	t.Run("overwrites HTTP 404 error messages", func(tt *testing.T) {
		e := &echo.Echo{}
		assert.Nil(tt, e.HTTPErrorHandler)
		RegisterErrorHandler(e, app)
		assert.NotNil(tt, e.HTTPErrorHandler)
	})
}
