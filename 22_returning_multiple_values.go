// Returnting Multiple Values
// - Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
// - Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya di function

// Menghiraukan/Ignore Return Value
// - Multiple return value wajib ditangkap semua valuenya
// - Jika kita ingin menghiraukan/ignore return value tersebut, kita bisa menggunakan tanda _ (underscore)

package main

import "fmt"

func getFullName() (string, string) {
  return "Rizqi", "Assyaufi"
}

func main() {
  firstName, lastName := getFullName()
  fmt.Println(firstName, lastName)

  fmt.Println(getFullName())

  // firstName = getFullName() // error
  firstName, _ = getFullName()
  fmt.Println(firstName)
}
