package main

import "fmt"

/* Interface Kosong */

// return value => any atau interface{}
func Ups(value any) interface{} {
  return value
}

func main() {
  var ups any = Ups("Rizqi")
  fmt.Println(ups)

  ups = Ups(123)
  fmt.Println(ups)

  ups = Ups(true)
  fmt.Println(ups)
}
