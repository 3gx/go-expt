package main

import "fmt"

func average(xs []float64) float64 {
  total := 0.0
  for _, v := range xs {
    total += v
  }
  return total / float64(len(xs))
}

func f() (int, int) {
  return 5, 6
}

func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

func makeEvenGenerator() func() uint {
  i := uint(0)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

func factorial(x uint) uint {
  if x == 0 {
    return 1
  }
  return x * factorial(x-1)
}

func first() {
  fmt.Println("1st")
}
func second() {
  fmt.Println("2nd")
}
func third() {
  fmt.Println("3rd")
}

func main() {
  {
    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs))

    someOtherName := []float64{198,193,177,182,183}
    fmt.Println(average(someOtherName))
  }
  {
    x, y := f()
    fmt.Println(x,y)
    fmt.Println(add(1,2,3))
    xs := []int{1,2,3}
    fmt.Println(add(xs...))
  }
  {
    add := func(x, y int) int {
      return x + y
    }
    fmt.Println(add(1,1))
    x := 0
    increment := func() int {
      x++
      return x
    }
    fmt.Println(increment())
    fmt.Println(increment())
  }
  {
    nextEven := makeEvenGenerator()
    fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) // 4
  }
  fmt.Println(factorial(7))
  {
    defer second()
    defer third()
    first()
  }
  {
    f := func() {
      defer third()
      defer second()
      first()
    }
    f()

  }
  {
    defer func() {
      str := recover()
      fmt.Println(str)
    }()
    panic("PANIC")
  }
}

