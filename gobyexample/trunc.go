package main

import "fmt"

func main() {
    // Variables
    var number int = 10
    message := "Go Programming Demo"

    fmt.Println("Message:", message)
    fmt.Println("Initial Number:", number)

    // Pointer
    ptr := &number
    *ptr = 20 // Modify value using pointer
    fmt.Println("Modified Number via Pointer:", number)

    // If-Else
    if number > 15 {
        fmt.Println("Number is greater than 15")
    } else {
        fmt.Println("Number is 15 or less")
    }

    // Switch
    switch number {
    case 10:
        fmt.Println("Number is 10")
    case 20:
        fmt.Println("Number is 20")
    default:
        fmt.Println("Number is something else")
    }

    // For Loop
    fmt.Println("Counting from 1 to 5:")
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }

    // NEW PART — Reading a floating-point number and truncating it
    var floatInput float64
    fmt.Print("Enter a floating point number: ")
    fmt.Scan(&floatInput)

    truncated := int(floatInput)
    fmt.Println("Truncated integer:", truncated)
}