package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// whole point is to return a db variable and let any file that uses the package
// interact with this
var DB *gorm.DB

func Connect() {
	// TODO: find the postgres credentials and create dsn string
	dsn := "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = d
}

func GetDB() *gorm.DB {
	return DB
}
