package main

import "fmt"

func main() {
  var (
    a = 5
    b = 2
    c = a + b
    d = a + b - c
    e = c / (a + b)
    f = a * b
    g = a + b - d * e / f
  )

  fmt.Println(c) // 7
  fmt.Println(d) // 0
  fmt.Println(e) // 1
  fmt.Println(f) // 10
  fmt.Println(g) // 7

  // augmented assignment
  a += 5 // a = a + 5
  b -= 1 // b = b - 1

  fmt.Println(a) // 10
  fmt.Println(b) // 1

  // unary operator
  a-- // a = a - 1
  b++ // b = b + 1

  fmt.Println(a) // 9
  fmt.Println(b) // 2
  fmt.Println(+a) // 9
  fmt.Println(-b) // -2
}
