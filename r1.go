package main
import "fmt"

func Recursive(number int) int {

    if number == 1 {

        return number
    }

    return number + Recursive(number-1)
}

func main() {

    answer := Recursive(2)
    fmt.Printf("Recursive: %d\n", answer)
}