package main

import (
    "fmt"
    "strings"
)

func main() {
    var input string
    fmt.Println("Enter a string:")
    fmt.Scanln(&input)

    // Convert to lowercase for case-insensitive check
    input = strings.ToLower(input)

    // Check if string contains 'i', 'a', and 'n' in order
    if strings.Contains(input, "i") && strings.Contains(input, "a") && strings.Contains(input, "n") {
        fmt.Println("Found!")
    } else {
        fmt.Println("Not Found!")
    }
}