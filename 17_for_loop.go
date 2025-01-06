package main

import "fmt"

func main() {
  counter := 1

  for counter <= 10 {
    fmt.Println("Perulangan ke", counter)
    counter++
  }

  // for with statement
  // we can add 2 statement:
  // - Init statement, yaitu statement sebelum for dieksekusi
  // - Post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan
  for counter = 1; counter <= 10; counter++ {
    fmt.Println("Perulangan ke", counter)
  }

  // for range
  // - for bisa digunakan untuk melakukan iterasi terhadap sema data collection
  // - data collection contohnya Array, Slice, dan Map
  names := []string{"Rizqi", "Nur", "Assyaufi"}

  for _, name := range names {
    fmt.Println(name)
  }

  for name := range names {
    fmt.Println(names[name])
  }

  for index, name := range names {
    fmt.Println("index", index, "=", name)
  }
}
