package main

import "fmt"

func main(){
	name := "haris"

	if name == "haris"{
		fmt.Println("Hello haris!")
	}else if name == "jacob" {
		fmt.Println("Hello jacob!")	
	}else {
		fmt.Println("Boleh kenalan gak??")
	}

	if length := len(name); length > 5 { // if short statement -> jadi dia mendeklarasikan variabel terlebih dahulu
		fmt.Println("Nama terlalu panjang")
	}else {
		fmt.Println("Nama sudah benar")
	}

	// switch length := len(name); length > 5 {
	// case true :
	// 	fmt.Println("Nama terlalu panjang")
	// case false: 
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)

	switch{
	case length >10 : 
		fmt.Println("Nama terlalu panjang")
	case name == "haris" :
		fmt.Println("Nama lumayan panjang")
	default : 
		fmt.Println("Nama sudah benar")
	}

}