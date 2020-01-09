package main

import "fmt"

func main() {
  {
    var x [5]int
    x[4] = 100
    fmt.Println(x)
  }

  {
    var x [5]float64
    x[0] = 98
    x[1] = 93
    x[2] = 77
    x[3] = 82
    x[4] = 83

    var total float64 = 0
    for i := 0; i < 5; i++ {
      total += x[i]
    }
    fmt.Println(total/5)

    total = 0;
    for _, value := range x {
      total += value
    }
    fmt.Println(total/float64(len(x)))

    y := [5]float64 {
      98,
      93,
      77,
      82,
      83,
    }
    total = 0;
    for _, value := range y {
      total += value
    }
    fmt.Println(total/float64(len(y)))
  }

  {
    var x []float64
    fmt.Println(len(x))
    y := make([]float64, 5)
    z := make([]float64, 5, 10)
    fmt.Println(len(y), len(z))
    arr := [5]float64{1,2,3,4,5}
    w := arr[1:3]
    fmt.Println(arr, w)
    arr[2]=44
    fmt.Println(arr, w)
  }

  {
    slice1 := []int{1,2,3}
    slice2 := append(slice1, 4,5)
    fmt.Println(slice1, slice2)
    slice1[1] = 33
    fmt.Println(slice1, slice2)
    slice2[1] = 7
    fmt.Println(slice1, slice2)
  }

  {
    x := make(map[string]int)
    x["key"] = 10
    fmt.Println(x)
    y := make(map[int]int)
    y[1] = 10
    y[2] = 10
    fmt.Println(y, y[1])
    delete(y,1)
    fmt.Println(y, y[2])
    delete(y,43)
  }

  {
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"

    fmt.Println(elements["Li"])

    {
      name, ok := elements["Li"]
      fmt.Println(name, ok)
    }
    {
      name, ok := elements["Un"]
      fmt.Println(name, ok)
    }
  }

  {
    elements := map[string]string{
      "H":  "Hydrogen",
      "He": "Helium",
      "Li": "Lithium",
      "Be": "Beryllium",
      "B":  "Boron",
      "C":  "Carbon",
      "N":  "Nitrogen",
      "O":  "Oxygen",
      "F":  "Fluorine",
      "Ne": "Neon",
    }
    name, ok := elements["Li"]
    fmt.Println(name, ok)
  }
  {
    {
      elements := map[string]map[string]string{
        "H": map[string]string{
          "name":"Hydrogen",
          "state":"gas",
        },
        "He": map[string]string{
          "name":"Helium",
          "state":"gas",
        },
        "Li": map[string]string{
          "name":"Lithium",
          "state":"solid",
        },
        "Be": map[string]string{
          "name":"Beryllium",
          "state":"solid",
        },
        "B":  map[string]string{
          "name":"Boron",
          "state":"solid",
        },
        "C":  map[string]string{
          "name":"Carbon",
          "state":"solid",
        },
        "N":  map[string]string{
          "name":"Nitrogen",
          "state":"gas",
        },
        "O":  map[string]string{
          "name":"Oxygen",
          "state":"gas",
        },
        "F":  map[string]string{
          "name":"Fluorine",
          "state":"gas",
        },
        "Ne":  map[string]string{
          "name":"Neon",
          "state":"gas",
        },
      }

      if el, ok := elements["Li"]; ok {
        fmt.Println(el["name"], el["state"])
      }
    }
  }
}
