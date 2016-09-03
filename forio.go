package main
import "fmt"

var xInt, testcase, num  int
var Ystr string
// var testcase int
// var num int

func Calculate(number int) int {
	
	// var xInt int
    str1 := 1
    // fmt.Println(str1)
    // fmt.Scanf("%d", &xInt)
    // fmt.Println("X Integer(s)", &xInt)
    return str1
}

func main() {
    
    str1 := "How many test case(s)?:"
    fmt.Print(str1)
    fmt.Scanf("%d\n", &testcase)
    fmt.Println("Testcase(s):",testcase)

    if testcase == 0 {
        fmt.Println("No test case...Do nothing")
    } else {
        for i := 0; i < testcase; i++ {
            // fmt.Println("Value of i is now:", i+1)
            fmt.Print("X Integer(s):")
            fmt.Scanf("%d\n", &xInt)
            fmt.Print("Y Integer(s):")
            fmt.Scanf("%s\n", &Ystr)
            fmt.Println("X Integer(s):", xInt)
            fmt.Println("Y Integer(s):", Ystr)

        }

    }
    
}