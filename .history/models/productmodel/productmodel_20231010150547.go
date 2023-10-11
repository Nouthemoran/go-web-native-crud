package productmodel

import "crud/entities"

func GetAll() []entities.Product {
	config.DB.Query(`
		SELECT
			products.id,
			products.name,
			categories.name as category_name
			products
	`)
}