package main

import "fmt"

func main() {
  // indirect declaration
  var names [3]string

  names[0] = "Rizqi"
  names[1] = "Nur"
  names[2] = "Assyaufi"

  fmt.Println(names[0])
  fmt.Println(names[1])
  fmt.Println(names[2])

  // direct declaration
  var fullName = [3]string{
    "Rizqi",
    "Nur",
    "Assyaufi",
  }

  fmt.Println(fullName[0], fullName[1], fullName[2]) // => Rizqi Nur Assyaufi

  var numbers1 = [3]int{1, 2, 3}
  numbers2 := [5]int{4, 5, 6, 7, 8}

  fmt.Println(numbers1) // => [1 2 3]
  fmt.Println(numbers2) // => [4 5 6 7 8]

  var values = [3]int{}
  fmt.Println(values) // => [0 0 0]

  // array's function
  // open length
  var numbers3 = [...]int{
    70,
    80,
    90,
    100,
    110,
  }

  fmt.Println(numbers3) // => [70, 80, 90, 100, 110]
  fmt.Println(len(numbers3)) // => 5
}
