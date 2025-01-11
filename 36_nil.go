package main

import "fmt"

/*
Nil
- Biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
- Berbeda dengan Go, di Go saat kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default value nya
- Namun di Go ada data nil, yaitu data kosong
- Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer, dan channel
*/

func NewMap(name string) map[string]string {
  if name == "" {
    return nil
  } else {
    return map[string]string{
      "name": name,
    }
  }
}

func main() {
  data := NewMap("")
  if data == nil {
    fmt.Println("Data map masih kosong")
  } else {
    fmt.Println(data)
  }
}
