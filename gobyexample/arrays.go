package main

import "fmt"

func main() {
	
	var a = []int{1, 2, 3, 4, 5}

	for i, v := range a {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	var b [5]int
	fmt.Println("emp:", b)
	
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	a = append(a, 6)
	a = append(a, 7, 8, 9)
	fmt.Println("append:", a)

	c := make([]int, len(a))
	copy(c, a)
	fmt.Println("copy:", c)

	l := a[2:5]
	fmt.Println("slice[2:5]:", l)

	l = a[:5]
	fmt.Println("slice[:5]:", l)

	l = a[2:]
	fmt.Println("slice[2:]:", l)

	// Multi-dimensional slices with make
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// Arrays
	b = [...]int{10, 20, 30, 40, 50}
	fmt.Println("array:", b)

	// Multi-dimensional array
	twoDArray := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2d array:", twoDArray)
}
