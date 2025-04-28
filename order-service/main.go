package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Order struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Item   string `json:"item"`
	User   *User  `json:"user,omitempty"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	orders []Order
	lock   = sync.Mutex{}
)

func createOrder(writer http.ResponseWriter, request *http.Request) {
	userID := request.FormValue("user_id")
	item := request.FormValue("item")

	if userID == "" || item == "" {
		http.Error(writer, "Missing user id or item", http.StatusBadRequest)
	}

	response, err := http.Get("http://localhost:8081/user?id=" + userID)

	if err != nil || response.StatusCode != http.StatusOK {
		http.Error(writer, "Invalid User", http.StatusBadRequest)
		return
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var user User

	err = json.Unmarshal(body, &user)

	if err != nil {
		http.Error(writer, "Error parsing User", http.StatusInternalServerError)
		return
	}

	lock.Lock()
	defer lock.Unlock()

	order := Order{
		ID:     len(orders) + 1,
		UserID: user.ID,
		Item:   item,
		User:   &user,
	}

	orders = append(orders, order)

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(orders)
	if err != nil {
		log.Println(err.Error())
		return
	}

}

func listOrders(writer http.ResponseWriter, request *http.Request) {
	lock.Lock()
	defer lock.Unlock()

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(orders)
	if err != nil {
		log.Println(err.Error())
		return
	}

}

func main() {
	http.HandleFunc("/orders", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {

		case http.MethodGet:
			listOrders(writer, request)

		case http.MethodPost:
			createOrder(writer, request)

		default:
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)

		}
	})

	fmt.Println("Order service listening on port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
