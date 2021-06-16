package server

import (
	"context"
	"fmt"
	"net/http"
	"test/api/handler/errors"
	"test/api/handler/health"
	"test/api/handler/user"
	"test/api/middleware/recovery"
	"test/api/middleware/secure"
	"test/internal/app"
	"test/pkg/signals"

	"github.com/labstack/echo"
)

func New(app app.App) *http.Server {
	e := echo.New()

	e.Use(secure.Middleware())
	e.Use(recovery.Middleware())
	e.Use(secure.MiddlewareBasicAuth(app.Config))

	health.RegisterRoutes(e)

	userGroup := e.Group("/v1/users")
	user.RegisterRoutes(userGroup, app)

	errors.RegisterErrorHandler(e, app)

	echo.NotFoundHandler = notFoundHandler

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Port),
		Handler: e,
	}

	graceful := signals.Setup()

	go func() {
		<-graceful
		err := srv.Shutdown(context.Background())
		if err != nil {
			app.Log.Error("Server ShutDown")
		}
	}()

	return srv
}

func notFoundHandler(c echo.Context) error {
	c.SetPath("/:path")
	return echo.NewHTTPError(http.StatusNotFound, "not found")
}
