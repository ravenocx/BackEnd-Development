package helper

import(
	"testing"
)

func BenchmarkHelloWorld(b *testing.B){
	// Penamaan func benchmark -> harus diawali dengan Benchmark dan terdapat parameter (b *testing.B)
	// nama file -> diakhiri _test (bisa digabung dengan unit test)
	// jika benchmark digabung dengan unit test, maka benchmark tidak akan di running

	// b.N sudah otomatis ditentukan oleh struct testing.B

	for i := 0; i < b.N ; i++ {
		HelloWorld("Haris")
	}

	// go test -v -bench=. -> running semua benchmark bersamaan dengan unit test pada satu package
	// go test -v -run=NotMathUnitTest -bench=. -> running benchmark tanpa unit test
	// NotMarhUnitTest -> nama func test yang tidak ada di file go
	// go test -v -run=NotMathUnitTest -bench=BenchmarkTest -> running specific benchmark
	// go test -v -bench=. ./... -> run semua benchmark di root module (folder paling atas)
}

