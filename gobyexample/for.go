package main

import "fmt"

func main() {
	// Using a labeled for loop to demonstrate breaking out of nested loops
	OuterLoop:
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				if i*j > 4 {
					break OuterLoop // Breaks out of the OuterLoop when condition is met
				}
				fmt.Printf("i: %d, j: %d, product: %d\n", i, j, i*j)
			}
		}
		fmt.Println("Exited the loops.")

	for i, value := range []int{10, 20, 30, 40, 50} {
		fmt.Printf("Index: %d , value: %d\n", i, value)
	}	

	for i := range 3 {
		fmt.Printf("This will not compile: Index: %d\n", i)		
	}

	for {
		fmt.Println("Infinite loop example. Breaking out immediately.")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			fmt.Printf("Even number: %d\n", n)
		} else {
			fmt.Printf("Odd number: %d\n", n)
		}
	}	
}
