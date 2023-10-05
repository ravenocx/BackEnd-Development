package helper

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M){
	// TestMain -> akan otomatis tereksekusi setiap menjalankan unit test di "sebuah package"
	// TestMain dieksekusi hanya sekali per Go-Lang package, bukan per tiap function unit test
	fmt.Println("Sebelum Unit Test")
	
	m.Run() // eksekusi semua unit test pada 1 package

	fmt.Println("Setelah Unit Test")
}