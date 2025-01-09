// Recover
// - Recover adalah function yang bisa kita gunakan untuk menangkap data panic
// - Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalanA

package main

import "fmt"

func endApp() {
  fmt.Println("End App")
  message := recover()
  fmt.Println("Terjadi Error:", message)
}

func runApp(error bool) {
  defer endApp()

  if error {
    panic("Ups, Error!")
  }

  // contoh penggunaan recover yang salah, karena tidak akan dieksekusi
  // message := recover()
  // fmt.Println("Terjadi Error:", message)
}

func main() {
  runApp(true)
  fmt.Println("Program Berakhir")
}
