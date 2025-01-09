// Defer
// - Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesi dieksekusi
// - Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

package main

import "fmt"

func logging() {
  fmt.Println("Selesai memanggil function")
}

func runApplication() {
  defer logging()
  fmt.Println("Run Application")
}

func main() {
  runApplication()
}
