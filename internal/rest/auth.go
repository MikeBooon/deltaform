package rest

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	mw "github.com/mikebooon/deltaform/internal/rest/middleware"
	"github.com/mikebooon/deltaform/service"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(e *echo.Echo, repo service.ServiceRepo) {
	handler := &AuthHandler{
		userService: repo.UserService,
	}

	e.POST("/auth/send-otp", handler.SendOTP, middleware.RateLimiter(mw.SecureRateLimitStore))
}

type SendOTPDTO struct {
	Email string `json:"email" validate:"required,email"`
}

func (h *AuthHandler) SendOTP(c echo.Context) error {
	var body SendOTPDTO

	err := c.Bind(&body)

	if err != nil {
		return err
	}

	err = c.Validate(body)

	if err != nil {
		return err
	}

	code, err := h.userService.GetNewVerificationCode(body.Email)

	if err != nil {
		log.Fatal(err)
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Failed to generate a one time code",
		)
	}

	log.Println(code)

	// NOW TO EMAIL CODE

	return c.NoContent(http.StatusAccepted)
}
