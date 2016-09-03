package main
import "fmt"

var xInt int
var testcase int
var num int

func Recursive(number int) int {
	
	// var xInt int
	// str1 := "How X Integer(s)?"
	// fmt.Println(str1)
	// fmt.Scanf("%d", &xInt)
	// fmt.Println("X Integer(s)", &xInt)

	fmt.Printf("Enter next number\n")
    n, err := fmt.Scanf("%d\n", &num)
    if err != nil {
        fmt.Println(n, err)
    }
    fmt.Println(num)

    if number == 1 {
        return number
    }
    return number + Recursive(number-1)
}

func main() {
    
    
    str1 := "How many test case(s)?"
    // xInt = 0
    fmt.Println(str1)
    fmt.Scanf("%d", &testcase)
    // fmt.Scanf("%s", &myname)
    // fmt.Println("Hello", testcase)

    // answer := Recursive(testcase)

    if testcase == 0 {
        fmt.Println("No test case...Do nothing")
    } else {
        answer := Recursive(testcase)
	    fmt.Printf("Recursive: %d\n", answer)
    }
    
}