package products

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func GetProductById(id string, db *sqlx.DB) (*Product, error) {
	var product Product
	fmt.Println(id)
	queryStatement := `SELECT productId, name, price FROM Products WHERE productId = ?;`

	err := db.Get(&product, queryStatement, id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(product)
	return &product, nil
}

func CreateNewProduct(p *Product, db *sqlx.DB) (*int64, error) {
	queryStatement := `INSERT INTO Products (name, price) VALUES (?, ?);`
	result, err := db.Exec(queryStatement, p.Name, p.Price)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	lastInserted, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &lastInserted, nil
}

func DeleteProductById(id string, db *sqlx.DB) error {
	queryStatement := `DELETE FROM Products WHERE (productId = ?);`
	_, err := db.Exec(queryStatement, id)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func UpdateProductById(id string, p *Product, db *sqlx.DB) (*int64, error) {
	queryStatement := `UPDATE Products SET name = ?, price = ? WHERE (productId = ?);`
	result, err := db.Exec(queryStatement, p.Name, p.Price, id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	lastInserted, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &lastInserted, nil

}
