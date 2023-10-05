package main

import "fmt"

func main(){
	fmt.Println(len("Hello World!"))
	fmt.Println("Hello World"[1])
	//-------------------------------------

	var umur int8 = 12

	var(
		firstName = "Haris"
		lastName = "Mulyadi"
	)

	const(
		friendName = "Alucard"
		friendName2 = "Jacob"
	)

	//-------------------------------------
	var value32 int32 = 128
	var value64 int64 = int64(value32)
	var value8 int8 = int8(value32) // value will be -127
	fmt.Println(umur,lastName,value64,value8)

	var h = firstName[0] //byte
	var hString = string(h) // string

	//-------------------------------------
	type identity string
	type married bool

	var noKTP identity = "123456789"
	var marriedStatus married = true

	fmt.Println(hString,noKTP,marriedStatus)
}