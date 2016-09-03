package main

import (
         // "bufio"
         "fmt"
         // "os"
 )

var xInt, testcase, num int

func Recursive(number int) int {
	
    if number == 1 {
        return number
    }
    return number + Recursive(number-1)
}


func Tcase(number int) int {
    
    //Getting how many X integers
    xInt = 1

    //Getting Y numberss    


    if number == 1 {
        return number
    }
    return number + Tcase(number-1)
}

func main() {

    // Getting how many test cases    
    str1 := "How many test case(s)?"
    fmt.Println(str1)
    fmt.Scanf("%d", &testcase)

    // verify testcase is a valid integer or not
    if testcase == 0 {
        fmt.Println("No test case...Do nothing")
    } else {
        answer := Tcase(testcase)
	    fmt.Printf("Recursive: %d\n", answer)
    }
    
}