// Named Return Values
// - Biasanya saat kita memberti tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function
// - Namun kita juga bisa membuat variable secara langsung di tipe data return function nya

package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
  firstName = "Rizqi"
  middleName = "Nur"
  lastName = "Assyaufi"

  return firstName, middleName, lastName
}

func main() {
  firstName, middleName, lastName := getCompleteName()

  fmt.Println(firstName, middleName, lastName)
}
