package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mikebooon/deltaform/internal/rest"
)

const (
	defaultAddress = ":3090"
)

func main() {
	log.Println("DeltaForm Starting...")

	e := echo.New()

	rest.NewFormHandler(e)

	e.Start(defaultAddress)
}
