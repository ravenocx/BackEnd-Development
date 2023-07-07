package main

import gethello "github.com/ritsuhaaa/BackEnd-Development/Go-Lang/src/Basics/Modul-17-AccessModifier/get-hello"

//import dengan nama yang berbeda dari nama folder

func main() {
	
	nama := gethello.Name
	gethello.SayHelloTo(nama)
}