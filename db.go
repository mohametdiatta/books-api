package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("books.db"))
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Book{})
	return db
}
