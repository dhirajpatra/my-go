// json.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User to hold the value
type User struct {
	Name string
	Age  int
}

func main() {
	// http handler function to decode json data
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "%s is %d years old!", user.Name, user.Age)
	})

	// http handler function to encode json data
	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Name: "John",
			Age:  25,
		}

		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":8080", nil)
}
