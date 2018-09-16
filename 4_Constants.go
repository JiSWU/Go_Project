package main

//example using constant

import "fmt"
import "math"

const s string = "constant"

func main() {
  fmt.Println(s)

  const n = 50000000000

  const d = 3e20/n
  fmt.Println(d)

  fmt.Println(int64(d))

  fmt.Println(math.Sin(n))

  const (
    Visa = "Visa"
    Master = "MasterCard"
    Amex = "American Express"
  )

  fmt.Println(Visa)
  fmt.Println(Master)
  fmt.Println(Amex)

}
