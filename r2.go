package main

import (
         // "bufio"
         "fmt"
         // "os"
 )

var xInt int
var testcase int
var num int

func Recursive(number int) int {
	
    if number == 1 {
        return number
    }
    return number + Recursive(number-1)
}



func RecursiveX(number int) int {
    
    // consolereader := bufio.NewReader(os.Stdin)
    // // fmt.Print("Enter X Integer(s) : ")

    // input, err := consolereader.ReadString('\n') // this will prompt the user for input

    // if err != nil {
    //     fmt.Println(err)
    //     os.Exit(1)
    // }

    // fmt.Println(input)


    fmt.Println("Number=",number)
    if number == 1 {
        return number
    }
    return number + RecursiveX(number-1)
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
        answer := RecursiveX(testcase)
	    fmt.Printf("Recursive: %d\n", answer)
    }
    
}