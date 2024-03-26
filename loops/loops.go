// Example of a for loop
package main

import "fmt"

func main() {
	// Standard for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Loop over elements in an array or slice
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Infinite loop with break statement
	count := 0
	for {
		fmt.Println("Infinite loop")
		count++
		// till count is 3
		if count == 3 {
			break
		}
	}

	// Loop with continue statement
	for i := 0; i < 5; i++ {
		// it will not print 2
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}
