// Function as Parameter
// - Function tidak hanya bisa kita simpan di dalam variable sebagai value
// - Namun juga bisa kita gunakan sebagai parameter untuk function lain

// Function Type Declaration
// - Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
// - Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter

package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilterType(name string, filter Filter) {
  filteredName := filter(name)
  fmt.Println("Hello,", filteredName)
}

func sayHeloWithFilter(name string, filter func(string) string) {
  filteredName := filter(name)
  fmt.Println("Hello,", filteredName)
}

func spamFilter(name string) string {
  if name == "Anjing" {
    return "***"
  } else {
    return name
  }
}

func main() {
  sayHeloWithFilter("Rizqi", spamFilter) // => Hello, Rizqi

  filter := spamFilter
  sayHeloWithFilter("Anjing", filter) // => Hello, ***

  sayHelloWithFilterType("Rizqi", spamFilter) // => Hello, Rizqi

  filter = spamFilter
  sayHelloWithFilterType("Anjing", filter) // => Hello, ***
}
