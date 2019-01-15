package main
import "fmt"

//Function definition
func plus(a int, b int) int {
    return a + b
}

//if types of parameters are same, It can be used as follows.
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {

    res := plus(1, 2) //call function
    fmt.Println("1+2 =", res)
    
    res = plusPlus(1, 2, 3)  //call function
    fmt.Println("1+2+3 =", res)
    
}
