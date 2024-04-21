package rest

import (
	"log"

	"github.com/labstack/echo/v4"
)

type FormHandler struct {
}

func NewFormHandler(e *echo.Echo) {
	handler := &FormHandler{}

	e.GET("/form", handler.FetchForm)
}

func (f *FormHandler) FetchForm(c echo.Context) error {
	log.Println("test fetch form")

	return c.JSON(200, "Success")
}
