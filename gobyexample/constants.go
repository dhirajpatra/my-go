package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

func main() {
	fmt.Println(Pi)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	const e = true
	fmt.Println(e)

	const s = "constant string"
	fmt.Println(s)

	var f = math.Sin(Pi / 4)
	fmt.Println(f)

	// Pi = 3.14159 // This will cause a compile-time error
}
