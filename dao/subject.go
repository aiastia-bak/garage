package dao

import (
	"garage/model"
	"log"
)

func GetSubject() {
	db := Database.Default
	var data model.Subject
	row := db.Where("id = '1'").Take(&data)
	if row.Error != nil {
		log.Println(row.Error)
	} else {
		log.Println(data)
	}
}
