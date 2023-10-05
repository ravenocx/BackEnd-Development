package main 

import "fmt"

func sayHelloTo(firstName string, secondName string){ //func multiple parameter
	fmt.Printf("Hello ", firstName, " and ", secondName)
}

// ------------------------------------------------------------------------------------------------------------
	
func getHello(name string) (greeting, fullName string) { //return named func
	greeting = "hello " + name
	fullName = "M Haris Sitompul"
	return  
}

// ------------------------------------------------------------------------------------------------------------

func getFullName()(string , string, string ){ // Multiple returned value
	return "Muhammad","Haris","Sitompul"
}

// ------------------------------------------------------------------------------------------------------------
	
//varagrs -> datanya bisa menerima lebih dari satu input , atau seperti array -> hanya bisa di parameter yg terakhir
// Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
// JIka parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma
	
func sumAll(number ...int)int{ //bisa memasukkan slice menjadi argument dalam variadic function
	total := 0
	for _, result := range number{ // varagrs akan otomatis mengubah dari argumen menjadi slice
		total += result 
	}
	return total
}

// ------------------------------------------------------------------------------------------------------------

func sayGoodBye(name string)string{
	return "Good Bye " + name
}

// ------------------------------------------------------------------------------------------------------------
	
func legalCheck(age int)string{ 
	if age >= 17 {
		return "You are legal!"
	}else {
		return "You are not legal!"
	}
}

// func legalAge(umur int, isLegal func(int) string){ 
// 	legal := isLegal(umur)
// 	fmt.Println("Hello", legal)
// }

type legalChecker func(int)string  //jika function terlalu panjang, bisa digunakan type alias

func legalAge(umur int, isLegal legalChecker){ //function as parameter
	legal := isLegal(umur)
	fmt.Println("Hello", legal)
}

// ------------------------------------------------------------------------------------------------------------

func registerUser(name string, blacklist func(string)bool){
	if blacklist(name){
		fmt.Println("You are blocked", name)
	}else{
		fmt.Println("Welcome", name)
	}
}

func main(){
	
	firstName,middleName, _  := getFullName() // _ -> return value tidak digunakan
	fmt.Println(firstName,middleName)

	// ------------------------------------------------------------------------------------------------------------

	numbers := []int{1,2,3,4,5,6,7}
	fmt.Println(sumAll(numbers...))

	// ------------------------------------------------------------------------------------------------------------

	hello := sayGoodBye // function as a value
	// tujuannya adalah agar bisa digunakan sebagai argument/parameter didalam function 
	fmt.Println(hello("Haris"))
	
	// ------------------------------------------------------------------------------------------------------------
	
	legalAge(19, legalCheck)

	// ------------------------------------------------------------------------------------------------------------
	blacklist := func(name string) bool{
		return name == "anjing"
	}

	registerUser("eko", blacklist)
	registerUser("anjing",func(name string) bool{
		return name == "anjing"
	})

	// ------------------------------------------------------------------------------------------------------------
	
	// Closure -> kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
	/*
	Closure adalah konsep dalam pemrograman yang mengacu pada kemampuan suatu fungsi untuk mengakses dan menyimpan 
	referensi ke variabel-variabel yang ada di lingkup di luar fungsi tersebut, 
	meskipun fungsi tersebut sudah selesai dieksekusi atau telah keluar dari lingkup asalnya.
	*/
	namaSaya := "Haris" 

	printNamaSaya := func() {
		namaSaya:= "Budi"
		fmt.Println(namaSaya)
	}

	fmt.Println(namaSaya)
	printNamaSaya() //closure, dimana budi bisa diakses walaupun scopenya masih didalam fungsi printNamaSaya

}