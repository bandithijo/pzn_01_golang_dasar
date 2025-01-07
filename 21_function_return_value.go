// Function Return Value
// - Function bisa mengembalikan data
// - Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function tersebut
// - Jika function tersebut kita deklatasikan dengan tipe data kembalian, maka kita wajib di dalam function nya kita harus mengembalikan data
// - Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return diikuti dengan datanya

package main

import "fmt"

func getHello(name string) string {
  hello := "Hello, " + name + "!"
  return hello
}

func main() {
  result := getHello("Rizqi")
  fmt.Println(result)

  fmt.Println(getHello("Go"))
}
