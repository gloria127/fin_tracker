package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func Connect(dbUrl string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		log.Fatalf("ошибка подключения к бд: %v\n", err)
	}
	log.Println("подключено к бд")
	return db
}
