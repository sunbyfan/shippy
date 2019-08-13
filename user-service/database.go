package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateConnection() (*gorm.DB, error) {
	host := "localhost"
	user := "postgres"
	DBName :="postgres"
	password :="123456"
	return gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, user, DBName, password,
		),
	)
}
