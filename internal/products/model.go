package products

type Product struct {
	Id    int     `db:"productId"`
	Name  string  `db:"name"`
	Price float64 `db:"price"`
}
