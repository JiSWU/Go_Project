package main

import "fmt"

func main() {

  var s = "initial"
  fmt.Println(s)

  var i,j int=1,2
  fmt.Println(i, j)

  var octal int = 0723       // Add '0'
  fmt.Println(octal)

  var hex int = 0x32fa2c75 // Add '0x
  fmt.Println(hex)

  var flag = true
  fmt.Println(flag)

  var novalue int
  fmt.Println(novalue)

  f := "short"          // use without 'var'
  fmt.Println(f)

}
