package models

import (
	"gorm.io/gorm"
)

type Books struct {
	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
	Author    *string `json: "author"`
	Titie     *string `json: "titie"`
	Publisher *string `json: "publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
