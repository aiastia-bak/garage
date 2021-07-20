package dao

import (
	"garage/model"
)

func GetAllJavTag() ([]model.JavTag, error) {
	var (
		db   = Database.Default
		data []model.JavTag
	)
	row := db.Model(&model.JavTag{}).Find(&data)
	if row.Error != nil {
		return nil, row.Error
	} else {
		return data, nil
	}
}

func CreateJavTag(data []model.JavTag) error {
	var db = Database.Default
	row := db.Create(&data)
	return row.Error
}
