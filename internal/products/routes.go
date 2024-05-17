package products

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupProdcutsRoutes(e *echo.Echo) {
	e.GET("/products/:id", getProducts)
	// e.POST("/products", saveProducts)
	// e.PUT("/products/:id", updateProducts)
	// e.DELETE("/products/:id", deleteUsers)
}

func getProducts(c echo.Context) error {
	id := c.Param("id")
	//threat info if needed

	//access controller passing id param

	return c.String(http.StatusOK, id)
}
