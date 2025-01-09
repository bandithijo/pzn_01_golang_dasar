package main

import "fmt"

/*
Struct Method
- Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
- Namun jika kita ingin menambahkan method ke dalam struct, sehingga seakan-akan sebuah struct memiliki function
- Method adalah function
*/

type Customer struct {
  Name, Address string
  Age           int
}

func (customer Customer) sayHello(name string) {
  fmt.Println("Hello", name ,"! my name is", customer.Name)
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

  // struct method
  customer1.sayHello("Jamal")
  customer2.sayHello("Jamal")
  customer3.sayHello("Jamal")
}
