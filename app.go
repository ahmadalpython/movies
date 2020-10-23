package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type App struct {
	DB *sql.DB
}

func (a *App) InitializeDB(host, user, password, dbname string) {
	args := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	var err error
	a.DB, err = sql.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database!")
}
