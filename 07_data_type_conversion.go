package main

import "fmt"

func main() {
  // convert integer
  var nilai32 int32 = 32768
  var nilai64 int64 = int64(nilai32)
  var nilai16 int16 = int16(nilai32)

  fmt.Println(nilai32) // => 32768
  fmt.Println(nilai64) // => 32768
  fmt.Println(nilai16) // => -32768 (number overflow)

  // convert string
  var name string = "Rizqi Nur Assyaufi"
  var n uint8 = name[6]
  var nString string = string(n)

  fmt.Println(name)
  fmt.Println(n)
  fmt.Println(nString)
}
