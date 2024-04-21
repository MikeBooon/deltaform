package rest

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikebooon/deltaform/service"
)

type FormHandler struct {
	formService service.FormService
}

func NewFormHandler(e *echo.Echo, repo service.ServiceRepo) {
	handler := &FormHandler{
		formService: repo.FormService,
	}

	e.GET("/form", handler.FetchForm)
}

func (f *FormHandler) FetchForm(c echo.Context) error {
	form, err := f.formService.GetByID(1)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	log.Println("test fetch form")

	return c.JSON(200, form)
}
