package main

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/Grypium/ansible-dashboard/handlers"
	"github.com/Grypium/ansible-dashboard/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

)

const (
	connStr = "user=your_user dbname=your_dbname password=your_password host=localhost sslmode=disable"
)

func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := models.InitDB(db); err != nil {
		log.Fatal(err)
	}

	// Log each request
	r.Use(handlers.Logging)

	// Users routes
	r.Handle("/users", handlers.Authenticate(handlers.WithDB(handlers.ListUsers, db))).Methods("GET")
	r.Handle("/users", handlers.Authenticate(handlers.WithDB(handlers.CreateUser, db))).Methods("POST")
	r.Handle("/users/{id}", handlers.Authenticate(handlers.WithDB(handlers.GetUser, db))).Methods("GET")
	r.Handle("/users/{id}", handlers.Authenticate(handlers.WithDB(handlers.UpdateUser, db))).Methods("PUT")
	r.Handle("/users/{id}", handlers.Authenticate(handlers.WithDB(handlers.DeleteUser, db))).Methods("DELETE")

	// Roles routes
	r.Handle("/roles", handlers.Authenticate(handlers.WithDB(handlers.ListRoles, db))).Methods("GET")
	r.Handle("/roles", handlers.Authenticate(handlers.WithDB(handlers.CreateRole, db))).Methods("POST")
	r.Handle("/roles/{id}", handlers.Authenticate(handlers.WithDB(handlers.GetRole, db))).Methods("GET")
	r.Handle("/roles/{id}", handlers.Authenticate(handlers.WithDB(handlers.UpdateRole, db))).Methods("PUT")
	r.Handle("/roles/{id}", handlers.Authenticate(handlers.WithDB(handlers.DeleteRole, db))).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
