package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users = []User{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Smith"},
	}
	lock = sync.Mutex{}
)

func getUserById(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")

	if id == "" {
		http.Error(writer, "ID is required", http.StatusBadRequest)
		return
	}

	lock.Lock()
	defer lock.Unlock()

	for _, user := range users {
		if fmt.Sprintf("%d", user.ID) == id {
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(user)
			return
		}
	}

	http.Error(writer, "User not found", http.StatusNotFound)
}

func listUsers(writer http.ResponseWriter, request *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(users)
}

func main() {
	http.HandleFunc("/users", listUsers)
	http.HandleFunc("/user", getUserById)

	fmt.Println("Listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
