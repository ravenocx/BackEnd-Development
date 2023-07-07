package main

import "fmt"

func endApp(){
	message := recover() 
	// recover -> function untuk menangkap data dari panic
	// dengan recover, proses panic akan terhenti, sehingga program akan tetap berjalan
	if message != nil {
		fmt.Println("Terjadi Error :", message)
	}
	fmt.Println("End app!")
}

func runApp (error bool){
	defer endApp()
	if error{
		panic("ERROR")
		// panic function -> dieksekusi ketika ada error dan akan langsung menghentikan program
		// saat panic function dipanggil, defer function tetap tereksekusi
		// ketika panic dipanggil, maka line selanjutnya diabaikan
	}
}

func main(){
	runApp(true)
}