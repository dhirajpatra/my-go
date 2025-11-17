package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Initialize an empty integer slice with an initial capacity of 3.
	// The length (size) is 0, but the underlying array has space for 3 elements
	// and will automatically grow as needed when 'append' is called.
	numbers := make([]int, 0, 3)

	// Set up a reader to handle console input
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Dynamic Sorted Integer Slice Program ---")
	fmt.Println("Enter an integer to add it to the sorted slice.")
	fmt.Println("Enter 'X' (or 'x') to quit the program.")

	// Start the main input loop
	for {
		fmt.Print("\nEnter an integer or 'X': ")

		// Read the user input until a newline character is encountered
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Clean up the input string (remove leading/trailing whitespace and newline)
		input = strings.TrimSpace(input)

		// Check the quit condition ('X' or 'x')
		if strings.EqualFold(input, "X") {
			fmt.Println("\nExiting program. Final sorted slice:", numbers)
			break
		}

		// Attempt to convert the cleaned input string to an integer
		num, err := strconv.Atoi(input)

		if err != nil {
			// Handle invalid input that is not 'X'
			fmt.Printf("Invalid input '%s'. Please enter a valid integer or 'X' to quit.\n", input)
			continue
		}

		// 1. Add the integer to the slice (the slice grows if capacity is exceeded)
		numbers = append(numbers, num)

		// 2. Sort the slice
		// sort.Ints sorts the slice in place in ascending order.
		sort.Ints(numbers)

		// 3. Print the sorted contents of the slice
		fmt.Println("Current sorted slice:", numbers)
	}
}