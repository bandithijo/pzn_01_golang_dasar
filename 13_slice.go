// Tipe Data Slice
// - Tipe data Slice adalah potongan dari data Array
// - Slice mrip dengan Array, yang membedakan ukuran Slice dapat berubah
// - Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array

// Detail Tipe Data Slice
// - Tipe data Slice memiliki 3 properties, yaitu: pointer, length, capacity
// - Pointer adalah penunjuk data pertama di Array pada Slice
// - Length adalah panjang dari Slice
// - Capacity adalah kapasitas dari Slice, dimana length tidak boleh lebih dari capacity

package main

import "fmt"

func main() {
  days := [...]string{
    "Senin",
    "Selasa",
    "Rabu",
    "Kamis",
    "Jumat",
    "Sabtu",
    "Minggu",
  }

  slice1 := days[3:6]
  fmt.Println(slice1) // => [Kamis Jumat Sabtu]

  var slice2 []string = days[5:]
  fmt.Println(slice2) // => [Sabtu, Minggu]

  // slice's functions
  fmt.Println(len(slice1)) // => 3
  fmt.Println(cap(slice1)) // => 4

  daySlice1 := days[5:]
  daySlice1[0] = "Sabtu Baru"
  daySlice1[1] = "Minggu Baru"
  fmt.Println(daySlice1) // => [Sabtu Baru Minggu Baru]
  fmt.Println(days) // => [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

  // append
  daySlice2 := append(daySlice1, "Libur Baru")
  fmt.Println(daySlice1) // => [Sabtu Baru Minggu Baru]
  fmt.Println(daySlice2) // => [Sabtu Baru Minggu Baru Libur Baru]
  daySlice2[0] = "Sabtu Lama"
  fmt.Println(daySlice2) // => [Sabtu Lama Minggu Baru Libur Baru]
  fmt.Println(days) // => [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

  // make slice
  newSlice := make([]string, 2, 5)
  fmt.Println(newSlice) // => [ ]
  newSlice[0] = "Rizqi"
  newSlice[1] = "Rizqi"
  fmt.Println(newSlice) // => [Rizqi Rizqi]
  fmt.Println(len(newSlice)) // => 2
  fmt.Println(cap(newSlice)) // => 5
  // newSlice[2] = "Rizqi" // error, harusnya menggunakan append
  newSlice2 := append(newSlice, "Rizqi")
  fmt.Println(newSlice2) // => [Rizqi Rizqi Rizqi]
  fmt.Println(len(newSlice2)) // => 3
  fmt.Println(cap(newSlice2)) // => 5
  newSlice[0] = "Budi"
  fmt.Println(newSlice2) // => [Budi Rizqi Rizqi]
  fmt.Println(newSlice) // =>[Budi Rizqi]

  // copy slice
  fromSlice := days[:]
  toSlice := make([]string, len(fromSlice), cap(fromSlice))
  fmt.Println(fromSlice) // => [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]
  fmt.Println(toSlice) // => [      ]
  copy(toSlice, fromSlice)
  fmt.Println(toSlice) // => [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

  // Hati-hati Saat Membuat Slice
  // - Saat membuat Array, kita harus berhati-hati, jika salah, maka yang kita buat bukanlah Array, melainkan Slice
  iniArray1 := [...]int{1, 2, 3, 4, 5}
  iniArray2 := [5]int{1, 2, 3, 4, 5}
  iniSlice := []int{1, 2, 3, 4, 5}
  fmt.Println(iniArray1) // => [1 2 3 4 5]
  fmt.Println(iniArray2) // => [1 2 3 4 5]
  fmt.Println(iniSlice) // => [1 2 3 4 5]
}
