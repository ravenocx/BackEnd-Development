package main

import "github.com/ritsuhaaa/BackEnd-Development/Go-Lang/src/Basics/Modul-16-Package-Import/helperr"

// file di golang hanya bisa mengakses file yang berada di package yang sama
// ketika ingin menggunakan file yang diluar package, bisa menggunakan import

// package -> sebuah direktori folder untuk mengorganisir kode program yang dibuat
// nama folder tidak boleh pake space

func main(){
	helperr.SayHello("Haris")
}
