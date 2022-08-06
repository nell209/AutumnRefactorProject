package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection(environment string) (*gorm.DB, error) {
	var dsn string
	if environment != "production" {
		dsn = "sslmode=disable port=5432 host=localhost user=postgres password=pa55word dbname=postgres"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
	//	 TODO add prod connection
}
