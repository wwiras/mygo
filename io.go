package main
import "fmt"

func main() {
    var myname string
    str1 := "What is your name?"
    fmt.Println(str1)
    // fmt.Sscan("%s\n", &myname)
    // fmt.Scanf("%s\n", &myname)
    fmt.Scanf("%s\n", &myname)
    // fmt.Scanln("%s\n", &myname)
    // fmt.Fscan("%s", &myname)
    // fmt.Scanln("%s", &myname)
    fmt.Println("Hello", myname)
}