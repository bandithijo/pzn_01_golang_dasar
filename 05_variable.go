package main

import "fmt"

func main() {
  // var name string
  // var name string = "Rizqi Nur"
  // var name = "Rizqi Nur"

  // name, age := "Rizqi Nur", 36

  name := "Rizqi Nur"
  age := 36
  fmt.Println(name, age)

  name = "Rizqi Nur Assyaufi"
  age = 25
  fmt.Println(name, age)

  // multiple declaration
  var (
    firstName = "Rizqi"
    middleName = "Nur"
    lastName = "Assyaufi"
  )
  fmt.Println(firstName, middleName, lastName)
}
