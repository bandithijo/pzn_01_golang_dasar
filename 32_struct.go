package main

import "fmt"

/*
Struct
- Astruct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
- Struct biasanya representasi data dalam program aplikasi yang kita buat
- Data di sturct disimpan dalam field/attribute
- Sederhananya struct adalah kumpulan dari field/attribute
- Nama struct dan nama attribute biasanya diawali huruf kapital

Membuat Data Struct
- Struct adalah template data atau prototype data
- Struct tidak bisa langsung digunakan
- Namun kita bisa membuat data/object dari struct yang telah kita buat

Struct Literals
- Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct
*/

type Customer struct {
  Name, Address string
  Age           int
}

func main() {
  var customer1 Customer

  fmt.Println(customer1) // => {  0}

  customer1.Name = "Rizqi Nur Assyaufi"
  customer1.Address = "Balikpapan, Kalimantan Timur"
  customer1.Age = 35

  fmt.Println(customer1) // => {Rizqi Nur Assyaufi Balikpapan, Kalimantan Timur 35}
  fmt.Println(customer1.Name) // => Rizqi Nur Assyaufi
  fmt.Println(customer1.Address) // => Balikpapan, Kalimantan Timur
  fmt.Println(customer1.Age) // => 35

  // struct literal
  customer2 := Customer{
    Name: "Devika Marina",
    Address: "Balikpapan, Kalimantan Timur",
    Age: 35,
  }
  fmt.Println(customer2) // => {Devika Marina Balikpapan, Kalimantan Timur 35}

  customer3 := Customer{"John Doe", "Auckland, New Zealand", 42}
  fmt.Println(customer3) // => {John Doe Auckland, New Zealand 42}
}
