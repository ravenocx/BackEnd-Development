package main

import "fmt"

func main(){
	for count :=1 ; count<=10 ;count++{// init;kondisi;post
		fmt.Printf("Perulangan ke ", count)
		// init statement -> sebelum for di eksekusi
		// post statement -> setelah for di eksekusi
	}

	slice := []string{"Muhammad", "Haris", "Sitompul"}

	// for i:= 0 ; i <= len(slice) ; i++ {
	// 	fmt.Println(slice[i])
	// }
	for index, name := range slice { 
		// index dari slice/string atau bisa juga key dari map
		// name -> value nya
		fmt.Printf("index", index, "=", name)
	}

	for _, name := range slice { 
		//underscore -> variable nya tidak dibutuhkan/digunakan
		fmt.Printf("name = %s", name)
	}
}