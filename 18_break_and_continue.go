// Break & Continue
// - Break * Continue adalah kata kunci yang bisa digunakan dalam perulangan
// - Break digunakan untuk menghentikan seluruh perulangan
// - Continue digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya
package main

import "fmt"

func main() {
  // break
  for i := 1; i <= 10; i++ {
    if i == 5 {
      break
    }

    fmt.Println("(BRK) Perulangan ke", i)
  }

  // continue
  for i := 1; i <= 10; i++ {
    if i % 2 == 0 {
      continue
    }

    fmt.Println("(CNT) Perulangan ke", i)
  }
}
