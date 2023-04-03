package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("Hello World!")

	var name string = "Haris"
	umur := 20

	//-------------------------------------
	identitas := "alucard_12022190353_si4501"

	upperString := strings.ToUpper(identitas)
	fmt.Println("Identity : ", upperString)

	//-------------------------------------
	fmt.Println("My name is %v", name)

	// fmt.Println("My name is %s and i'm %d years old", name, umur)
	fmt.Printf("My name is %s and I'm %d years old \n", name, umur)

	jumlah := 7.5
	fmt.Printf("There are %f kg apple in the table\n", jumlah)

	//-------------------------------------
	price := 16.540

	word := fmt.Sprintf("%s buy orange with price : $%.3f", name , price) // SprintF -> format a string into a variable
	fmt.Println(word)

	fmt.Printf("%0.f\n", price) // Dibulatkan keatas
    fmt.Printf("%0.5f\n", 16.540)

	//-------------------------------------
	n1 := 2
    n2 := 3
    n3 := 4

    res := fmt.Sprintf("There are %d oranges %d apples %d plums", n1, n2, n3)
    fmt.Println(res)

    res2 := fmt.Sprintf("There are %[2]d oranges %d apples %[1]d plums", n1, n2, n3)
    // fmt.Println(res2)
}