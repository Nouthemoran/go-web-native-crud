package categorymodel

import "crud/config"

func GetAll() {
	config.DB.Query(`SELECT * FROM categorymodel WHERE`) 
}