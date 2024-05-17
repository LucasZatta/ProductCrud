package products

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context, db *sqlx.DB) {
	id := c.Param("id")

	product, err := GetProductById(id, db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, &product)
}

func CreateProduct(c echo.Context, db *sqlx.DB) {
	price, err := strconv.ParseFloat(c.FormValue("price"), 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	p := Product{
		Name:  c.FormValue("name"),
		Price: price,
	}

	lastUpdatedId, err := CreateNewProduct(&p, db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, fmt.Sprintf("LastUpdatedId: %d", *lastUpdatedId))
}

func DeleteProduct(c echo.Context, db *sqlx.DB) {
	id := c.Param("id")

	err := DeleteProductById(id, db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Product with id %s was removed successfully!", id))
}

func UpdateProduct(c echo.Context, db *sqlx.DB) {
	id := c.Param("id")

	price, err := strconv.ParseFloat(c.FormValue("price"), 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	p := Product{
		Name:  c.FormValue("name"),
		Price: price,
	}

	lastUpdatedId, err := UpdateProductById(id, &p, db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, fmt.Sprintf("LastUpdatedId: %d", *lastUpdatedId))
}
