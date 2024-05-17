package main

import "github.com/LucasZatta/ProductCrud/internal/routes"

func main() {
	routes := routes.NewRouter()
	routes.Logger.Fatal(routes.Start(":1323"))

}
