package categorymodel

import "crud/config"

func GetAll() {
	rowsconfig.DB.Query(`SELECT * FROM categories`) 
}