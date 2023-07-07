package main

import (
	// _ -> bisa digunakan ketika kita hanya ingin menjalankan init function di package tanpa
	// harus mengeksekusi salah satu function atau variable yang ada di package
	_ "fmt"

	_"github.com/ritsuhaaa/BackEnd-Development/Go-Lang/src/Basics/Modul-18-PackageInit/database"
)

func main() {
	// fmt.Println(database.GetDatabase())
}