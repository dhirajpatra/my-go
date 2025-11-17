package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// The main function executes the logic for collecting user input,
// creating a map, converting it to JSON, and printing the result.
func main() {
	// 1. Get user input for Name
	var name string
	fmt.Print("Enter name: ")
	// Read the input line. If there is an error (like EOF), handle it.
	_, err := fmt.Scanln(&name)
	if err != nil && err.Error() != "unexpected newline" {
		fmt.Fprintf(os.Stderr, "Error reading name: %v\n", err)
		return
	}

	// 2. Get user input for Address
	var address string
	fmt.Print("Enter address: ")
	// Read the input line.
	_, err = fmt.Scanln(&address)
	if err != nil && err.Error() != "unexpected newline" {
		fmt.Fprintf(os.Stderr, "Error reading address: %v\n", err)
		return
	}

	// 3. Create a map and add the name and address
	// We use a map[string]string to store key-value pairs.
	dataMap := make(map[string]string)
	dataMap["name"] = name
	dataMap["address"] = address

	// 4. Use json.Marshal() to create a JSON object from the map
	// Marshal returns a byte slice and an error.
	jsonBytes, err := json.Marshal(dataMap)
	if err != nil {
		// Handle any error during the JSON marshaling process
		fmt.Printf("Error marshaling map to JSON: %v\n", err)
		return
	}

	// 5. Print the resulting JSON object
	// Convert the byte slice to a string for printing.
	fmt.Println("\n--- JSON Output ---")
	fmt.Println(string(jsonBytes))
	fmt.Println("-------------------")
}