// Anonymous Function
// - Sbeleunya setiap membuat function. kita akan selalu memberikan sebuah nama pada function tersebut
// - Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membua function terlebih dahulu
// - Hal tersebut dinamakan anonymous function

package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
  if blacklist(name) {
    fmt.Println("You are blocked,", name)
  } else {
    fmt.Println("Welcome,", name)
  }
}

func main() {
  // declare outside function call
  blacklist := func(name string) bool {
    return name == "anjing"
  }

  registerUser("Rizqi", blacklist) // => Welcome, Rizqi

  // declare inside function call
  registerUser("anjing", func(name string) bool {
    return name == "anjing"
  }) // => You are blocked, anjing
}
