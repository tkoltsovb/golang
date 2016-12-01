package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)

    var f = "short"
    fmt.Println(f)

    var s string = "s"
    fmt.Println(s)

    var a[5] int
    a[0] = 10
    fmt.Println(a[0])

    sl := make([]string, 3)
    sl[1] = "rt"
    fmt.Println("emp:", sl)
}