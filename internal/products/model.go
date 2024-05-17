package products

type Product struct {
	Id    int     `db:"idProducts"`
	Name  string  `db:"name"`
	Price float64 `db:"price"`
}
