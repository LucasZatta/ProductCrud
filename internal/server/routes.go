package server

import (
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
	e.DELETE("/products/:id", s.DeleteProduct)
	e.PUT("/products/:id", s.UpdateProduct)

	return e
}

func (s *Server) GetProducts(c echo.Context) error {
	products.GetProducts(c, s.db.ReturnDbInstance())
	return nil
}

func (s *Server) SaveProducts(c echo.Context) error {
	products.CreateProduct(c, s.db.ReturnDbInstance())

	return nil
}

func (s *Server) DeleteProduct(c echo.Context) error {
	products.DeleteProduct(c, s.db.ReturnDbInstance())

	return nil
}

func (s *Server) UpdateProduct(c echo.Context) error {
	products.UpdateProduct(c, s.db.ReturnDbInstance())

	return nil
}
