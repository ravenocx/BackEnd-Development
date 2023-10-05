package helper

import (
	"fmt"
	"testing"
)


func TestHelloWorldFailHaris(t *testing.T){
	result := HelloWorldFail("Haris")
	if result != "Hello Haris"{
		t.Error("Result must be 'Hello Haris'")
	}
	fmt.Println("TestHelloWorldFailHaris Done")
}
// t.Fail() -> menggagalkan unit test, namun tetap melanjutkan eksekusi kode program selanjutnya
// t.FailNow() -> menggagalkan unit test, tanpa melanjutkan eksekusi kode program selanjutnya
// t.Error(args...) -> log print error dan memanggil func Fail()
// t.Fatal(args...) -> log print error dan memanggil func FailNow()

func TestHelloWorldFailDimas(t *testing.T){
	result := HelloWorldFail("Dimas")
	if result != "Hello Dimas"{
		t.Fatal("Result must be 'Hello Dimas'")
	}
	fmt.Println("TestHelloWorldFailDimas Done")
}

