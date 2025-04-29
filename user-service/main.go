package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type User struct {
	ID   int    `json:"id"` // <- this is the struct field tag for [de]serialization
	Name string `json:"name"`
}

var (
	users = []User{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Smith"},
	}
	mutex = sync.Mutex{}
)

func getUserById(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")

	if id == "" {
		http.Error(writer, "ID is required", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for _, user := range users {
		if fmt.Sprintf("%d", user.ID) == id {
			writer.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(writer).Encode(user)
			if err != nil {
				log.Println(err.Error())
				return
			}

			return
		}
	}

	http.Error(writer, "User not found", http.StatusNotFound)
}

func listUsers(writer http.ResponseWriter, request *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(users)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func main() {
	http.HandleFunc("/users", listUsers)
	http.HandleFunc("/user", getUserById)

	fmt.Println("Listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
