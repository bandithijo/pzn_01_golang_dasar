package main

import "fmt"

/*
Pointer

Pass by Value
- Secara default di Go, semua variable itu di passing by value, bukan by reference
- Artinya, jika kita mengirim sebuah variable ke dalam function, method, atau variable lain, sebenarnya yang dikirim adalah duplikasi value nya

Pointer
- Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
- Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

Operator &
- Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator &, diikuti dengan nama variable nya
*/

type Address struct {
  City, Province, Country string
}

func main() {
  // default behaviour
  address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
  address2 := address1 // copy value

  address2.City = "Bandung"

  fmt.Println(address1) // => {Subang Jawa Barat Indonesia}
  fmt.Println(address2) // => {Bandung Jawa Barat Indonesia}

  // pointer
  address3 := Address{"Balikpapan", "Kalimantan Timur", "Indonesia"}
  address4 := &address3 // pass by reference

  address4.City = "Samarinda"

  fmt.Println(address3) // => {Samarinda Kalimantan Timur Indonesia}
  fmt.Println(address4) // => &{Samarinda Kalimantan Timur Indonesia}
}
