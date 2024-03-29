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

// Check if database exists
func CheckDBExists(db *sql.DB, name string) (bool, error) {
	var result string
	err := db.QueryRow("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?", name).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // Database doesn't exist
		}
		return false, err // Other error occurred
	}
	return true, nil // Database exists
}

func main() {
	// Configure database connection (always check error)
	// need to check the container internal IP with docker inspect command
	time.Sleep(5 * time.Second) // Wait for 5 seconds

	db, err := Connect("root", "pass", "localhost", "") // Connect without specifying database
	if err != nil {
		panic(err.Error())
	}

	// Initialize the first connection to the database to check if everything is working fine
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database connection established!")

	// Check if database exists
	exists, err := CheckDBExists(db, "test")
	if err != nil {
		panic(err.Error())
	}

	// If database exists, drop it
	if exists {
		err = DropDB(db, "test")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Database dropped!")
	}
	err = CreateDB(db, "test")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database created!")

	// Select the database
	_, err = db.Exec("USE test")
	if err != nil {
		panic(err.Error())
	}

	// create a table
	query := `create table employees (
        id int,
        name varchar(255),
        age int,
        salary int)`
	err = CreateTable(db, query)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table created!")

	query = `insert into employees (id, name, age, salary) values (1, 'John', 30, 5000)`
	_, err = db.Exec(query)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Data inserted!")

	query = `CREATE TABLE users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	err = CreateTable(db, query)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table created!")

	// create new user
	username := "dhirajpatra"
	password := "password"
	createdAt := time.Now()

	// insert user into database
	result, err := db.Exec("INSERT INTO users (username, password, createdAt) VALUES (?, ?, ?)", username, password, createdAt)
	if err != nil {
		panic(err.Error())
	}

	userID, err := result.LastInsertId()

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("User inserted! New user id is ", userID)
	}

	// call all users
	type user struct {
		id        int
		username  string
		password  string
		createdAt string
	}

	rows, err := db.Query("SELECT id, username, password, createdAt FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			panic(err.Error())
		}
		// Parse the createdAt string from the database into a time.Time object.
		// The layout parameter "2006-01-02 15:04:05" defines the expected format of the createdAt string.
		// The layout follows the reference time "Mon Jan 2 15:04:05 MST 2006", as specified in the Go documentation.
		// The values in the layout represent the year (2006), month (01), day (02), hour (15, in 24-hour format), minute (04), second (05).
		// The time.Parse function returns a time.Time object representing the parsed time.
		// If the parsing fails due to an invalid format, it returns an error.
		createdAtTime, err := time.Parse("2006-01-02 15:04:05", u.createdAt)

		if err != nil {
			panic(err.Error())
		}
		u.createdAt = createdAtTime.String() // Convert back to string if needed
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	fmt.Println("All users:")
	for _, u := range users {
		fmt.Printf("ID: %d, Username: %s, Password: %s, Created At: %s\n", u.id, u.username, u.password, u.createdAt)
	}

	// delete a user
	_, err = db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("User deleted!")

	db.Close()
}
