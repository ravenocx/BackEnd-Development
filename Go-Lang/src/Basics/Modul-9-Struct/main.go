package main

import "fmt"

type Customer struct {
	// struct -> representasi data dalam program (template data/prototype data)
	// data di struct disimpan dalam field (collection of fields)
	// nama struct dan field harus diawali dengan huruf besar
	Name, Address string //field
	Age	int
}

func (customer Customer) sayHello(){ // struct method -> berguna ketika diimplementasikan dengan interface
	fmt.Println("Hello my name is", customer.Name)
}

func main(){
	var haris Customer // Mengakses struct harus lewat objek
	haris.Name = "Muhammad Haris"
	haris.Address = "Jakarta"
	haris.Age = 20

	fmt.Println(haris)

	joko:= Customer {
		Name : "Joko", 
		Address : "Jakarta",
		Age : 20,
	}

	// budi:= Customer{"Budi", "Indonesia",30} //tidak direkomendasikan
	joko.sayHello()
}