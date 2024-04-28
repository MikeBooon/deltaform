package middleware

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mikebooon/deltaform/domain"
	"github.com/mikebooon/deltaform/internal/auth"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return echo.NewHTTPError(401, "Unauthorized")
		}

		if len(strings.Split(authHeader, " ")) != 2 {
			return echo.NewHTTPError(401, "Unauthorized")
		}

		token := strings.Split(authHeader, " ")[1]

		user, err := auth.ValidateJWT(token)

		if err != nil {
			log.Println(err.Error())
			return echo.NewHTTPError(401, "Unauthorized")
		}

		domainUser := domain.User{
			ID:    user.ID,
			Email: user.Email,
		}

		c.Set("user", domainUser)

		return next(c)
	}
}
