package main

import "fmt"

func printNumber1() {
  var number1 int
  number1 = 1
  fmt.Println(number1)
}

func printNumber2() {
  var number1 int = 2
  fmt.Println(number1)
}

func printNumber3() {
  number1 := 3
  fmt.Println(number1)
}

func printNumber4() {
  const number1 = 4
  fmt.Println(number1)
}

func printArray1() {
  arr1 := [3]string{}

  arr1[0] = "1"
  arr1[1] = "2"
  arr1[2] = "3"

  fmt.Println(arr1)
}

func printArray2() {
  arr1 := [10]string{}

  arr1[0] = "1"
  arr1[1] = "2"
  arr1[2] = "3"
  arr1[3] = "4"

  fmt.Println(arr1)
}

func printArray3() {
  arr1 := [...]int{1, 2, 3}
  fmt.Println(arr1)
}

func printLenCap() {
  arr1 := [...]int{1, 2, 3}
  fmt.Println(len(arr1))
  fmt.Println(cap(arr1))
}

func printSlice() {
  arr1 := []int{}
  arr1 = append(arr1, 1)
  arr1 = append(arr1, 2)
  arr1 = append(arr1, 3)

  fmt.Println(arr1)
}

func printMap() {
  map1 := map[string]int{
    "x": 100,
    "y": 200,
  }

  map1["z"] = 300
  fmt.Println(map1)

  delete(map1, "z")
  fmt.Println(map1)

  key := "key"
  map1[key] = 400
  fmt.Println(map1)

  fmt.Println(map1["x"])
  fmt.Println(map1["y"])
  fmt.Println(map1[key])

  _, isExistsKey := map1[key]

  // fmt.Println(_) underscore cant not defined
  fmt.Println(isExistsKey)

  for key, value := range map1 {
    fmt.Printf("%s = %d\n", key, value)
  }
}

func printIf () {
  const x = 100
  if x == 200 {
    fmt.Println("should not print")
  } else if x == 300 {
    fmt.Println("should not print")
  } else {
    fmt.Println("should print")
  }
}

func printSwitch() {
  mode := "shouldPrint"

  switch mode {
    case "shouldPrint":
      fmt.Println("should print")
      fallthrough // breakはないので、明示的次に行く
    case "shouldPrint2":
      fmt.Println("should print")
    case "shouldNotPrint":
      fmt.Println("should not print")
    default:
      fmt.Println("should not print")
  }
}

func printFor() {
  x := 0
  y := 2
  for x < y {
    x++
    fmt.Println(x)
  }

  for i := 0; i < 2; i++ {
    fmt.Println(i)
  }
}

func main() {
  printNumber1()
  printNumber2()
  printNumber3()
  printNumber4()

  printArray1()
  printArray2()
  printArray3()

  printLenCap()
  printSlice()
  printMap()

  printIf()
  printSwitch()

  printFor()
}
