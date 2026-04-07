package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Room struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Capacity int    `json:"capacity"`
}

func main() {

	connStr := "user=egorgoloseevgmail.com dbname=booking sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/rooms", getRooms)

	log.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}

func getRooms(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT id, name, location, capacity FROM rooms")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()

	rooms := []Room{}

	for rows.Next() {
		var r Room
		rows.Scan(&r.ID, &r.Name, &r.Location, &r.Capacity)
		rooms = append(rooms, r)
	}

	json.NewEncoder(w).Encode(rooms)
	http.Handle("/", http.FileServer(http.Dir("./static")))
}
