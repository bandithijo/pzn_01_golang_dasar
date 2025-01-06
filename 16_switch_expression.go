package main

import "fmt"

func main() {
  name := "Rizqi"

  switch name {
  case "Rizqi":
    fmt.Println("Hello, Rizqi!")
  case "Budi":
    fmt.Println("Hello, Budi!")
  default:
    fmt.Println("Hi, Boleh kenalan?")
  }

  // switch short statement
  switch length := len(name); length > 5 {
  case true:
    fmt.Println("Nama Terlalu Panjang")
  case false:
    fmt.Println("Nama Sudah Benar")
  }

  // switch without first declaration
  nilai := 51
  switch {
  case nilai >= 90:
    fmt.Println("Nilai kamu: A+")
  case nilai >= 80:
    fmt.Println("Nilai kamu: A")
  case nilai >= 70:
    fmt.Println("Nilai kamu: B")
  case nilai >= 60:
    fmt.Println("Nilai kamu: C")
  case nilai >= 50:
    fmt.Println("Nilai kamu: D")
  default:
    fmt.Println("Nilai kamu: E")
  }
}
