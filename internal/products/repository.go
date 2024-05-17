package products

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func FindProductById(id string, db *sqlx.DB) {
	var product Product
	fmt.Println(id)
	queryStatement := `SELECT idProducts, name FROM Products WHERE idProducts = ?;`

	err := db.Get(&product, queryStatement, id)

	// result := db.QueryRow(queryStatement, id)
	// err := result.Scan(product)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(product)
}

func CreateProduct() {

}
