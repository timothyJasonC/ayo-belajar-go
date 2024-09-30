package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct untuk mewakili data user
type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

// Handler untuk menangani request ke endpoint /users
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Contoh data statis yang akan dikirim ke front-end
	users := []User{
		{Name: "John Doe", Age: 30, Address: "123 Main St"},
		{Name: "Jane Smith", Age: 25, Address: "456 Oak St"},
	}

	// Set header response sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode data ke JSON dan kirimkan response
	json.NewEncoder(w).Encode(users)
}

func main() {
	// Menangani endpoint /users
	http.HandleFunc("/users", getUsersHandler)

	// Menjalankan server di port 8000
	fmt.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
