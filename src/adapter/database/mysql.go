package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySql struct {
	*gorm.DB
}

func SetupMysql() *MySql {
	fmt.Println("Starting to connect database")
	dsn := "root@tcp(localhost:3306)/db_crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &MySql{db}
}
