package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Connect establishes a database connection.
func Connect(username, password, host, database string) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, host, database)
	db, err := sql.Open("mysql", conn)
	return db, err
}

// Employee struct
type Employee struct {
	ID      int
	Name    string
	Salary  int
	City    string
	State   string
	Country string
}

// EmployeePageData contains the data needed to render the employee list page.
type EmployeePageData struct {
	PageTitle string
	Employees []Employee
}

func main() {
	// Configure database connection
	time.Sleep(5 * time.Second) // Wait for 5 seconds

	// my-mysql container connection details
	db, err := Connect("root", "pass", "localhost", "test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Initialize the first connection to the database to check if everything is working fine
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database connection established!")

	// Select the database
	_, err = db.Exec("USE test")
	if err != nil {
		panic(err.Error())
	}

	query := `SELECT e.id, e.name, e.salary, a.city, a.state, a.country 
	FROM employees e 
	INNER JOIN employee_addresses a ON a.employee_id = e.id`

	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var e Employee
		// read the values from the current row in the result set and store them into the specified variables reference
		err := rows.Scan(&e.ID, &e.Name, &e.Salary, &e.City, &e.State, &e.Country)
		if err != nil {
			panic(err.Error())
		}
		employees = append(employees, e)
	}
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	fmt.Println("All employees:")
	for _, e := range employees {
		fmt.Printf("ID: %d, Name: %s, Salary: %d, City: %s, State %s, Country %s\n", e.ID, e.Name, e.Salary, e.City, e.State, e.Country)
	}

	// Parse the HTML template file.
	tmpl := template.Must(template.ParseFiles("employee.html"))

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define the handler for the root URL.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create the data to be passed to the template.
		data := EmployeePageData{
			PageTitle: "My Employee list",
			Employees: employees,
		}
		// Execute the template with the data and write the result to the response.
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server and listen for incoming requests.
	fmt.Println("Server started on port 80")
	http.ListenAndServe(":80", nil)
}
