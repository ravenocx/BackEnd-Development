package gethello 

import "fmt"

var Name = "Haris" // Upper case harus menggunakan var =


// nama function atau variable yang diawali dengan huruf besar -> bisa diakses oleh package lain

func SayHelloTo(name string){
	fmt.Println("Hello", name)
}