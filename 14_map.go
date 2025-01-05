// Tipe Data Map
// - Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
// - Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
// - Sederhananya, Map adalah tipe data kumpulan key-value, dimana kata kuncinya bersifat unik, tidak boleh sama
// - Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kuncinya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

package main

import "fmt"

func main() {
  // var person map[string]string = map[string]string{}
  // fmt.Println(person) // => map[]
  // person["name"] = "Rizqi"
  // person["address"] = "Balikpapan, Kalimantan Timur"

  person := map[string]string{
    "name": "Rizqi",
    "address": "Balikpapan, Kalimantan Timur",
  }

  // person := make(map[string]string)

  fmt.Println(person) // => map[address:Balikpapan, Kalimantan Timur name:Rizqi]
  fmt.Println(person["name"]) // => Rizqi
  fmt.Println(person["address"]) // => Balikpapan, Kalimantan Timur

  person["name"] = "Nur"
  fmt.Println(person) // => map[address:Balikpapan, Kalimantan Timur name:Nur]
  fmt.Println(person["name"]) // => Nur

  // map's function
  // len(map) | Untuk mendapatkan jumlah data di map
  // map[key] | Mengambil data di map dengan key
  // map[key] = value | Mengubah data di map dengan key
  // make(map[TypeKey]TypeValue) | Membuat map baru
  // delete(map, key) | Menghapus data di map dengan key

  book := make(map[string]string)
  book["title"] = "Buku Golang"
  book["author"] = "Rizqi Assyaufi"
  book["publish"] = "2025"

  fmt.Println(book) // => map[author:Rizqi Assyaufi publish:2025 title:Buku Golang]

  delete(book, "publish")
  fmt.Println(book) // => map[author:Rizqi Assyaufi title:Buku Golang]
}
