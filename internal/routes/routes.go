package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/products/:id", getProducts)
	// e.POST("/products", saveProducts)
	// e.PUT("/products/:id", updateProducts)
	// e.DELETE("/products/:id", deleteUsers)

	return e
}

func getProducts(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
