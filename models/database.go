package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	if err := db.AutoMigrate(&Requests{}); err != nil {
		log.Println(err)
	}

	return db, err
}
