package rest

import (
	"net/http"
	"strconv"

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

	e.GET("/form/:id", handler.get)
}

func (f *FormHandler) get(c echo.Context) error {
	idStr := c.Param("id")

	if idStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing param: id")
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid param: id")
	}

	form, err := f.formService.GetByID(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	return c.JSON(200, form)
}
