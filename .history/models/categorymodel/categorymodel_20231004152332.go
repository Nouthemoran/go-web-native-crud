package categorymodel

import "crud/config"

func GetAll() {
	rows.config.DB.Query(`SELECT * FROM categories`) 
}