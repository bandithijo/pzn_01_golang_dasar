package main

import "fmt"

/*
Type Assertions
- Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
- Fitur ini sering sekali digunakan ketika kita bertemu dengan data interafce kosong

Type Assertions Menggunakan Switch
- Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
- Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
- Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
*/

func random(value any) interface{} {
  return value
}

func main() {
  var result any = random(123)

  // var resultString string = result.(string)
  // fmt.Println(resultString)

  // var resultInt int = result.(int) // panic
  // fmt.Println(resultInt)

  switch value := result.(type) {
  case string:
    fmt.Println("String", value)
  case int:
    fmt.Println("Int", value)
  default:
    fmt.Println("Unknown", value)
  }
}
