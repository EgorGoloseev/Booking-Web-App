package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var db *sqlx.DB

type User struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Role string `db:"role" json:"role"`
}

type Room struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Location string `db:"location" json:"location"`
	Capacity int    `db:"capacity" json:"capacity"`
}

type Booking struct {
	ID        int       `db:"id" json:"id"`
	RoomID    int       `db:"room_id" json:"room_id"`
	UserID    int       `db:"user_id" json:"user_id"`
	StartTime time.Time `db:"start_time" json:"start_time"`
	EndTime   time.Time `db:"end_time" json:"end_time"`
	Purpose   string    `db:"purpose" json:"purpose"`
}

func main() {

	conn := "postgres://booking:booking@localhost:5432/booking?sslmode=disable"

	var err error

	db, err = sqlx.Connect("postgres", conn)

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/rooms", roomsHandler)
	http.HandleFunc("/bookings", bookingHandler)

	log.Println("server started :8080")

	http.ListenAndServe(":8080", nil)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	users := []User{}

	err := db.Select(&users, "SELECT * FROM users")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func roomsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		rooms := []Room{}

		err := db.Select(&rooms, "SELECT * FROM rooms")

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(rooms)
	}

	if r.Method == "POST" {

		var room Room

		json.NewDecoder(r.Body).Decode(&room)

		query := `
		INSERT INTO rooms(name,location,capacity)
		VALUES($1,$2,$3)
		`

		_, err := db.Exec(query,
			room.Name,
			room.Location,
			room.Capacity,
		)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(201)
	}
}

func bookingHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		var booking Booking

		json.NewDecoder(r.Body).Decode(&booking)

		if booking.EndTime.Before(booking.StartTime) {

			http.Error(w, "end_time > start_time", 400)

			return
		}

		query := `
		SELECT COUNT(*)
		FROM bookings
		WHERE room_id=$1
		AND NOT ($3 <= start_time OR $2 >= end_time)
		`

		var count int

		db.Get(&count,
			query,
			booking.RoomID,
			booking.StartTime,
			booking.EndTime)

		if count > 0 {

			http.Error(w, "time conflict", 400)

			return
		}

		insert := `
		INSERT INTO bookings
		(room_id,user_id,start_time,end_time,purpose)
		VALUES($1,$2,$3,$4,$5)
		`

		_, err := db.Exec(insert,
			booking.RoomID,
			booking.UserID,
			booking.StartTime,
			booking.EndTime,
			booking.Purpose)

		if err != nil {

			http.Error(w, err.Error(), 500)

			return
		}

		w.WriteHeader(201)
	}
}

conn := "postgres://booking:booking@localhost:5432/booking"

db, err := sqlx.Connect("postgres", conn)