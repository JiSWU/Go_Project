package main
import "fmt"

//This function returns 2 ints.
func vals() (int, int) {
    return 3, 7
}

func main() {

    a, b := vals() //return 2 ints
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals() // return 1 int
    fmt.Println(c)
}
