package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Connect establishes a database connection.
func Connect(username, password, host, database string) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, host, database)
	db, err := sql.Open("mysql", conn)
	return db, err
}

// CreateDB creates a new database.
func CreateDB(db *sql.DB, name string) error {
	query := fmt.Sprintf("CREATE DATABASE %v", name)
	_, err := db.Exec(query)
	return err
}

// CreateTable creates a new table in the database.
func CreateTable(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	return err
}

// DropDB drops an existing database.
func DropDB(db *sql.DB, name string) error {
	query := fmt.Sprintf("DROP DATABASE %v", name)
	_, err := db.Exec(query)
	return err
}

func main() {
	// Configure database connection (always check error)
	// need to check the container internal IP with docker inspect command
	time.Sleep(5 * time.Second) // Wait for 5 seconds

	db, err := Connect("root", "pass", "localhost", "test")
	if err != nil {
		panic(err.Error())
	}

	// Initialize the first connection to the database to check if everything is working fine
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database connection established!")
}
