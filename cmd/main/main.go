package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//// DATABASE
		//dbSettings := psql.Settings{
		//	User:   "krakra",
		//	Pass:   "krapassword",
		//	Host:   "localhost",
		//	Port:   "5432",
		//	Name:   "kra_db",
		//	Reload: false,
		//}
		//err := psql.Connect(dbSettings)
		//if err != nil {
		//	log.Fatal(err)
		//}

		db, err := sqlx.Connect("postgres", "user=t_user dbname=t_db sslmode=disable password=456654 host=localhost")
		if err != nil {
			log.Fatalln(err)
		}

		defer db.Close()

		// Test the connection to the database
		var user User
		if err := db.Ping(); err != nil {
			user = User{
				Name:  "John failed Doe",
				Email: "john@example.com",
			}
		} else {
			user = User{
				Name:  "John success Doe",
				Email: "john@example.com",
			}
		}

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
