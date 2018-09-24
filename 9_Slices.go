package main
import "fmt"
func main() {

    s := make([]string, 3)
    fmt.Println("emp:", s)  //empty array

    s[0] = "a" //input 'a'
    s[1] = "b" //input 'b'
    s[2] = "c" //input 'c'
    fmt.Println("set:", s) //print index from 0 to 2.
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    //add 'd' to the end
    s = append(s, "d") 
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
