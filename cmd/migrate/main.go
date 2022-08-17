package main

import (
	"database/sql"
	"fmt"
	_ "github.com/WHGLD/antibruteforce/migrations"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/pressly/goose"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	dbConnection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Name,
	)

	err = goose.SetDialect("postgres")
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := sql.Open("pgx", dbConnection)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err := goose.Run("up", db, "./migrations/"); err != nil {
		fmt.Println(err)
		return
	}
}
