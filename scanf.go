package main

import (
    "fmt"
)

// compute square roots by using Newton's method
func main() {
    var x float64 //number to take square root
    var y float64 //this is the guess
    var q float64 //this is the quotient
    var a float64 //this is the average

    fmt.Print("Enter a number to take its square root: ")
    var inputSquare float64
    n, err := fmt.Scanf("%f\n", &inputSquare)
    if err != nil || n != 1 {
        // handle invalid input
        fmt.Println(n, err)
        return
    }

    fmt.Print("Enter first guess ")
    var inputGuess float64
    n, err = fmt.Scanf("%f\n", &inputGuess)
    if err != nil || n != 1 {
        // handle invalid input
        fmt.Println(n, err)
        return
    }

    x = inputSquare
    y = inputGuess
    for i := 0; i < 10; i++ {
        q = x / y       // compute the quotient; x and y are given
        a = (q + y) / x // compute the average
        y = a           // set the guess to the average
    }
    fmt.Printf("sqrt(%g) = %g\n", x, y)
}