package main

import (
	"fmt"
	"github.com/ritsuhaaa/go-say-hello"
)

func main(){
	// release versi module -> git tag v1.0.0 -> git push origin v1.0.0
	// menambah dependency -> go get nama-module 

	// untuk upgrade module -> melakukan git tag lagi
	// upgrade dependency -> mengubah tag di go mod (v....) -> go get

	// jika ada major update -> tambahkan /v2 dibelakangnya

	// go get -> untuk mengunduh dan menginstal paket atau modul tertentu 
	// dari repositori publik atau pribadi ke dalam workspace
	// go mod tidy ->  memperbarui file go.mod dan go.sum Anda berdasarkan 
	// dependensi yang sebenarnya digunakan dalam kode sumber
	
	// ketika ada package yang sama tetapi berbeda versi (major update v1->v2)
	// go get -> menginstall versi yang digunakan dengan tidak menghapus versi yang lama apabila prefix berbeda
	// tetapi go get apabila prefixnya sama maka hanya di upgrade v nya saja
	// go mod tidy -> menginstall versi yang digunakan dan menghapus package yang tidak digunakan di go.mod dan go.sum

	result := gosayhello.SayHello()
	fmt.Println(result)
}