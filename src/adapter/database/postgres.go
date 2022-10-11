package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	*gorm.DB
}

func SetupPostgres() *PostgreSQL {
	fmt.Println("Starting to connect database")
	dsn := "host=localhost user=postgres password=postgres dbname=db_retail port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &PostgreSQL{db}
}
