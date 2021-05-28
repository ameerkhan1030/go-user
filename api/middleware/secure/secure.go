package secure

import (
	"test/pkg/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Middleware returns an Echo middleware function that adds HSTS-related headers
// to all responses.
func Middleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			res := c.Response()

			res.Header().Set(echo.HeaderXXSSProtection, "1; mode=block")
			res.Header().Set(echo.HeaderXContentTypeOptions, "nosniff")
			res.Header().Set("X-Download-Options", "noopen")
			res.Header().Set("Referrer-Policy", "same-origin")
			res.Header().Set(echo.HeaderStrictTransportSecurity, "max-age=31536000; includeSubdomains; preload")
			res.Header().Set(echo.HeaderContentSecurityPolicy, "default-src 'none'; object-src 'self'; block-all-mixed-content")

			return next(c)
		}
	}
}

// MiddlewareBasicAuth return an middleware function that performs basic auth verification.
func MiddlewareBasicAuth(config config.Config) func(next echo.HandlerFunc) echo.HandlerFunc {
	authConfig := middleware.BasicAuthConfig{
		Validator: func(username, password string, c echo.Context) (bool, error) {
			if config.BasicAuthUsername == username && config.BasicAuthPassword == password {
				return true, nil
			}
			return false, nil
		},
		Skipper: func(c echo.Context) bool {
			if c.Request().URL.RequestURI() == "/health" {
				return true
			}
			return false
		},
	}
	return middleware.BasicAuthWithConfig(authConfig)
}