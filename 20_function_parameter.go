// Function Parameter
// - Saat membuat function, kadang-kadang kita membutuhkan data dari luar, atau kita sebut dengan parameter
// - Kita bisa menambahkan parameter di function, bisa lebih dari satu
// - Parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat
// - Namun jika kita menambahkan parameter di function, maka ketika kita memanggil function tersebut, kita wajib memasukkan data ke parameternya

package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
  fmt.Println("Hello,", firstName, lastName)
}

func main() {
  sayHelloTo("Rizqi", "Assyaufi")
}
