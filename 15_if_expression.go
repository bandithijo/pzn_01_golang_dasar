package main

import "fmt"

func main() {
  name := "Rizqi"

  if name == "Rizqi" {
    fmt.Println("Hello, Rizqi!")
  } else if name == "Budi" {
    fmt.Println("Hello, Budi!")
  } else if name == "Joko" {
    fmt.Println("Hello, Joko!")
  } else if name == "Rizqi" {
    fmt.Println("Hello, Rizqi Lagi!") // tidak akan pernah dieksekusi
  } else {
    fmt.Println("Hi, Boleh kenalan?")
  }

  if name == "Rizqi" {
    fmt.Println("Hello,", name)
  }

  // if with short statement
  // only if statement using by those condition
  if length := len(name); length > 5 {
    fmt.Println("Nama terlalu panjang")
  } else {
    fmt.Println("Nama sudah benar")
  }
}
