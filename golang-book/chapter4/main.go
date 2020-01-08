package main

import "fmt"

func main() {
  {
    var x string = "Hello World"
    fmt.Println(x)
  }
  {
    var x string
    x = "Hello World"
    fmt.Println(x)
  }
  {
    var x string
    x = "first"
    fmt.Println(x)
    x = "second"
    fmt.Println(x)
  }
  {
    var x string
    x = "first"
    fmt.Println(x)
    x = x + "second"
    fmt.Println(x)
  }
  {
    var x string = "hello"
    var y string = "world"
    fmt.Println(x == y)
  }
  {
    x := "Hello world"
    fmt.Println(x)
    y := 5
    fmt.Println(y)
  }
}

