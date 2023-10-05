package main

import "fmt"

type Address struct {
	City,Provincy,Country string
}

func changeCountryToIndonesia(address Address){ 
	address.Country = "Indonesia"
}

func ChangeCountryToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func(address *Address) AddKota(){ // di struct method juga sama pass by value
	address.City = "Kota " + address.City
}

func main(){
	address1 := Address{"Subang","Jawa Barat",""}
	changeCountryToIndonesia(address1)
	fmt.Println(address1) // tidak berubah krn parameter di function itu pass by value

	ChangeCountryToIndonesia(&address1)
	fmt.Println(address1)

	address1.AddKota()
	fmt.Println(address1)


}