package main

import (
	"net/http"

	_c "github.com/developer-unkowns-code-ground/ez-backend/category"
	"github.com/labstack/echo"
)

func main() {
	// Inisialisasi Echo Server
	e := echo.New()

	_c.NewHttpCategoryHandler(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
