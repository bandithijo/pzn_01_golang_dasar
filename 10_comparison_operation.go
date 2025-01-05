package main

import "fmt"

func main() {
  var (
    name1 = "Rizqi"
    name2 = "Rizqi"

    result bool = name1 == name2
  )

  fmt.Println(result) // true

  result = name1 != name2
  fmt.Println(result) // false

  var (
    angka1 = 10
    angka2 = 20
  )

  fmt.Println(angka1 > angka2) // false

  result = angka1 < angka2
  fmt.Println(result) // true

  result = angka1 != angka2
  fmt.Println(result) // true

  var (
    alphabet1 = "a"
    alphabet2 = "b"
  )

  result = alphabet1 > alphabet2
  fmt.Println(result) // false

  result = alphabet1 < alphabet2
  fmt.Println(result) // true
}
