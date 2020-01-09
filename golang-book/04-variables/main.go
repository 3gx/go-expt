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
  {
    var(
      a = 5
      b = 10
      c = 15
    )
    fmt.Println(a,b,c)
  }
  {
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)
    output := input * 2
    fmt.Println(output)
  }
}

