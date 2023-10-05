package helper

import "testing"

// aturan nama file test -> akhiran_test , ex : hello_world_test
// nama function untuk unti test -> harus diawali dengan Test -> TestHelloWorld
// dan harus memiliki parameter (t *testing.T) dan tidak mengembalikan return value

func TestHelloWorld(t *testing.T){
	/*
	The *testing.T type is a pointer to the testing.T struct, which holds information 
	about a specific test and provides methods for running assertions and logging test output.
	*/
	result := HelloWorld("Haris")
	if result != "Hello Haris"{
		panic("Result is not Hello Haris")
	}
}

// go test -> run semua unit test yang ada di dalam module atau package 
// go test -v -> melihat function apa saja yang di run ketika testing pada 1 package
// go test -v -run TestNamaFunction -> running test pada satu function yang kita inginkan
// go test ./... -> run semua unit test di root module (folder paling atas)

