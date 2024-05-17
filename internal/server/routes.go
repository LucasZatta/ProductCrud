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

	products.SetupProdcutsRoutes(e)

	return e
}
