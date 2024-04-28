package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mikebooon/deltaform/domain"
	"github.com/mikebooon/deltaform/internal/rest/middleware"
	"github.com/mikebooon/deltaform/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(e *echo.Echo, repo service.ServiceRepo) {
	handler := &UserHandler{
		userService: repo.UserService,
	}

	e.GET("/user/me", handler.Me, middleware.IsAuthenticated)
}

func (h *UserHandler) Me(c echo.Context) error {
	user := c.Get("user").(domain.User)

	if user == (domain.User{}) {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	return c.JSON(200, user)
}
