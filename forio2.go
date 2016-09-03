package main

import "fmt"

func main() {
    fmt.Print("Enter text: ")
    var input string
    n, err := fmt.Scanln(&input) 
    fmt.Println(n, err, input) 
    // fmt.Print(input)
}