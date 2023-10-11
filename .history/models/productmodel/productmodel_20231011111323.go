package productmodel

import (
	"crud/config"
	"crud/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT
			products.id,
			products.name,
			categories.name as category_name,
			products.stock,
			products.description,
			products.created_at,
			products.updated_at
		FROM products
		JOIN categories ON products.category_id = categories.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.Stock,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

func Create(product entities.Product) bool {
	config.DB.Exec(
		INSERT INTO products(name)
	)
}
