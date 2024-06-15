package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	dsn := "host=localhost user=root password=root dbname=root port=54320 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Could not open database")
	}
}
