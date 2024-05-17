package server

import (
	"fmt"
	"net/http"

	"github.com/LucasZatta/ProductCrud/internal/products"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/products/:id", s.GetProducts)
	e.POST("/products", s.SaveProducts)

	// e.PUT("/products/:id", updateProducts)
	// e.DELETE("/products/:id", deleteUsers)
	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) GetProducts(c echo.Context) error {
	id := c.Param("id")
	products.FindProductById(id, s.db.ReturnDbInstance())
	return nil
}
func (s *Server) SaveProducts(c echo.Context) error {
	name := c.FormValue("name")
	price := c.FormValue("price")
	fmt.Println(name)
	fmt.Println(price)

	return nil
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
