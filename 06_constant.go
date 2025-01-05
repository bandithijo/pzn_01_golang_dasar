package main

import "fmt"

func main() {
  const name string = "Rizqi Nur"
  const age = 36
  fmt.Println(name, age)

  // error
  // name = "Rizqi Nur Assyaufi"
  // age = 25
  // fmt.Println(name, age)

  // multiple declaration
  const (
    firstName = "Rizqi"
    middleName = "Nur"
    lastName = "Assyaufi"
  )
  fmt.Println(firstName, middleName, lastName)
}
