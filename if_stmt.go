package main

import "fmt"

func main(){
    a := 100
    
    fmt.Printf("a ")
    if a > 100{
        fmt.Printf("bigger than 100\n")
    } else if a == 100{
        fmt.Printf("is 100\n")
    } else if a < 100 {
        fmt.Printf("less than 100\n")
    }
}
