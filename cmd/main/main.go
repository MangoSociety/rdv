package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, World!")
	})

	err := http.ListenAndServe(":5555", nil)
	if err != nil {
		fmt.Println("error starting server: ", err)
	}
}
