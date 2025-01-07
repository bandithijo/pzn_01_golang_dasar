// Variadic Function
// - Parameter yang berada di posisi terakhir, memiliki kemampuan untuk dijadikan sebuah varargs (variable arguments)
// - Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array
// - Apa bedanya dengan parameter biasa dengan tipe data Array?
//   - Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
//   - Jika parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma

// Slice Parameter dalam Variadic Function

package main

import "fmt"

func sumAll(numbers ...int) int {
  total := 0

  for _, number := range numbers {
    total += number
  }

  return total
}

func sumAllSlice(numbers []int) int {
  return sumAll(numbers...)
}

func sumAllVariadic(numbers ...int) int {
  return sumAll(numbers...)
}

func main() {
  // sum with slice
  totalSlice := sumAllSlice([]int{1, 2, 3, 4, 5})
  fmt.Println(totalSlice) // => 15
  fmt.Println(sumAllSlice([]int{10, 10, 10, 10, 10, 10, 10})) // => 70

  // sum with variadic parameter
  totalVariadic := sumAllVariadic(1, 2, 3, 4, 5)
  fmt.Println(totalVariadic) // => 15
  fmt.Println(sumAllVariadic(10, 10, 10, 10, 10, 10, 10)) // => 70

  // slice argument inside variadic parameter
  numbers := []int{1, 2, 3, 4, 5}
  totalVariadic = sumAllVariadic(numbers...)
  fmt.Println(totalVariadic) // => 15
}
