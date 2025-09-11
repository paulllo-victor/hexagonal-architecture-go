package main

import (
	"database/sql"
	"log"

	db2 "github.com/paulllo-victor/go-hexagonal/adapters/db"
	"github.com/paulllo-victor/go-hexagonal/application"
)

func main() {

	db, _ := sql.Open("sqlite3", "sqlite.db")

	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, err := productService.Create("Product 1", 25)

	if err != nil {
		log.Fatal(err.Error())
	}

	productService.Enable(product)

}
