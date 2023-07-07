package main

import "fmt"

type HasName interface{ // hasName -> kontrak
	// interface hanya berisikan definisi-definisi method
	// interface digunakan sebagai kontrak
	GetName() string
}

func SayHello(hasName HasName){ 
	fmt.Println("Hello", hasName.GetName())
}

type Customer struct{
	Name string
}

func (customer Customer) GetName()string { // struct method
	return customer.Name
}

func Ups(apapun interface{})interface{}{
	// interface kosong -> tidak ada deklarasi method satupun
	// hal ini membuat secara otomatis semua tipe data akan menjadi impelementasinya (seperti parent class/super class)
	// interface kosong juga bisa dijadikan parameter ketika ingin tipe data yang masuk beragam
	if apapun == 1 {
		return 1
	}else if apapun == "haris" {
		return true
	} else {
		return "Upsie"
	} // returnya bebas tipe datanya karena secara tidak langsung mengikuti kontrak interface kosong
}

func main(){
	haris:= Customer{"Harissss"}
	SayHello(haris)
	/*
	cara memasukkan parameter di say hello adalah membuat struct 
	yang memiliki method yang sama  dengan method (kontrak) yang ada di interface.
	Karena struct yang memiliki nama method yang sama dengan kontrak di interface ,
	akan otomatis dianggap sebagai type interface tersebut
	*/

	data := Ups(2) // interface kosong
	fmt.Println(data)

	//interface kosong memiliki kemampuan untuk mewakili tipe data apa pun
	var emptyInterface interface{} // Deklarasi variabel dengan tipe interface kosong

	emptyInterface = 5      // Variabel emptyInterface dapat menampung nilai dengan tipe data apapun
	fmt.Println(emptyInterface)

	emptyInterface = "Hello" // Bisa juga diisi dengan tipe data string
	fmt.Println(emptyInterface)

	emptyInterface = true    // Bisa juga diisi dengan tipe data boolean
	fmt.Println(emptyInterface)
}