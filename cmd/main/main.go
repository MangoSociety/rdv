package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		db, err := sqlx.Connect("postgres", "user=t_user dbname=t_db sslmode=disable password=456654 host=localhost")
		if err != nil {
			log.Fatalln(err)
		}

		defer db.Close()

		// Test the connection to the database
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		} else {
			log.Println("Successfully Connected")
		}

		// Создание объекта пользователя
		user := User{
			Name:  "John Doe",
			Email: "john@example.com",
		}

		// Установка заголовка Content-Type
		w.Header().Set("Content-Type", "application/json")

		// Кодирование объекта пользователя в JSON и запись ответа
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	err := http.ListenAndServe(":5555", nil)
	if err != nil {
		fmt.Println("error starting server: ", err)
	}
}
