package main

import "fmt"

func main() {
    
    fmt.Println("this is i")
    i := 1
    for i<=3{
        fmt.Println(i)
        i++
    }
    
    fmt.Println("this is j")
    for j:=7; j<=9; j++{
        fmt.Println(j)
    }
    
    for{
        fmt.Println("for loop")
        break
    }
    
    for n:=0; n<=5; n++{
        if n%2 == 0{
            continue //if continue, don't execute fmt.Println(n)
        }
        fmt.Println(n)
    }
}
