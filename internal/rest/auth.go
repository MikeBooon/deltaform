package rest

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mikebooon/deltaform/internal/email"
	mw "github.com/mikebooon/deltaform/internal/rest/middleware"
	"github.com/mikebooon/deltaform/service"
)

type AuthHandler struct {
	userService service.UserService
	emailClient email.EmailClient
}

func NewAuthHandler(e *echo.Echo, repo service.ServiceRepo, emailClient email.EmailClient) {
	handler := &AuthHandler{
		userService: repo.UserService,
		emailClient: emailClient,
	}

	e.POST("/auth/send-otp", handler.sendOTP, middleware.RateLimiter(mw.SecureRateLimitStore))
	e.POST("/auth/verify-otp", handler.verifyOTP, middleware.RateLimiter(mw.SecureRateLimitStore))
}

type SendOTPDTO struct {
	Email string `json:"email" validate:"required,email"`
}

func (h *AuthHandler) sendOTP(c echo.Context) error {
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

	log.Println(code)

	if err != nil {
		log.Fatal(err)
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Failed to generate a one time code",
		)
	}

	err = h.emailClient.SendOTPEmail(body.Email, code)

	if err != nil {
		log.Fatal("Failed to send otp code email", err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusAccepted)
}

type VerifyOTPDTO struct {
	Code  string `json:"code" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (h *AuthHandler) verifyOTP(c echo.Context) error {
	var body VerifyOTPDTO

	err := c.Bind(&body)

	if err != nil {
		return err
	}

	err = c.Validate(body)

	if err != nil {
		return err
	}

	valid, user, err := h.userService.VerifyVerficationCode(body.Code, body.Email)

	if err != nil {
		log.Println(err)
		return c.NoContent(
			http.StatusInternalServerError,
		)
	}

	if !valid {
		return c.NoContent(
			http.StatusUnauthorized,
		)
	}

	token, err := h.userService.NewUserJWT(user.Email, user.ID)

	if err != nil {
		log.Fatal(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
