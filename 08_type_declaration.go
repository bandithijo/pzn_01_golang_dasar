package main

import "fmt"

func main() {
  // tipe data yang bernama NoKTP, yang merupakan alias dari string
  type NoKTP string

  var ktpRizqi NoKTP = "1111111111"

  var contoh string = "2222222222"
  var contohKtp NoKTP = NoKTP(contoh)

  fmt.Println(ktpRizqi)
  fmt.Println(contohKtp)
}
