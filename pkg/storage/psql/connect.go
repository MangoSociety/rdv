package psql

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose"
	"log"
)

type Settings struct {
	Host   string
	Port   string
	Name   string
	User   string
	Pass   string
	Reload bool
}

func Connect(settings Settings) (err error) {

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s, dbname=%s sslmode=disable",
		settings.Host, settings.Port, settings.User, settings.Pass, settings.Name)

	DB, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		return err
	}
	log.Printf("Database connection was created: %s \n", sqlInfo)

	if settings.Reload {
		log.Printf("Start reloading database \n")
		err := goose.DownTo(DB, ".", 0)
		if err != nil {
			return err
		}
	}
	// TODO: Settings migrations
	log.Printf("Start migrating database \n")
	err = goose.Up(DB, "migrations")
	if err != nil {
		return err
	}

	return nil
}
