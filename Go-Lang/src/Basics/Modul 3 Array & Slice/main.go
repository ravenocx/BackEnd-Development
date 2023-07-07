package main 

import "fmt"

func main(){
	// Array -> tipe data sama & jumlahnya tetap (tidak bisa ditambah/dikurangi)
	fruits := [6]string{"apple", "grape", "banana", "melon"}
	// fruit2 := [...]string{"apple", "banana", "melon"} // jumlah array akan disesuaikan dengan jumlah elemen yang ada
	// var friendName [2]string

	fmt.Println(fruits)

	// Slice -> potongan dari data array / data yang mengakses sebagian atau seluruh data di array
	// Slice tidak memiliki data, hanya referensi ke data asli (slice dan array selalu terkoneksi)
	// Slice mirip dengan array dalam cara mengaksesnya dan ukuran slice bisa berubah
	sliceFruits := fruits[1:4]
	fmt.Println(sliceFruits)
	apppendSliceFruits := append(sliceFruits, "kiwi") 
	//membuat slice baru , jika kapasitas slice sudah penuh maka akan membuat array baru

	apppendSliceFruits[1] = "papaya"
	fmt.Println(apppendSliceFruits)
	fmt.Println(sliceFruits)
	fmt.Println(fruits)

	sliceFruits2 := fruits[1:5]
	fmt.Println(sliceFruits2)


	/*
	array[low:] -> slice dari index low sampai index terakhir
	array[:high] -> slice dari index 0 sampai "high-1"
	array[:] -> slice dari index 0 sampai index terakhir
	pointer -> penunjuk data pertama di array pada slice
	length -> panjang slice (jumlah data di slice) -> (index terakhir -1 ) - index pertama
	capacity -> kapasitas slice (jumlah data yang bisa ditampung pada slice) -> jumlah data di array - index pertama di slice
	ex : friendName [12]string{"Haris", "Alucard", "Jacob", "Lancelot", "Gusion", "Layla", "Miya", "Eudora", "Nana", "Lolita", "Bruno", "Roger"}
	friendNameSlice := friendName[4:7]
	pointer = 4
	length = 3
	capacity = 12 - 4 = 8
	*/

	newSlice := make([]string, 2, 5) // make(type data, length, capacity)
	var newSlice2 []string // default nya 0
	// otomatis terbuat array dengan panjang 5
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(len(newSlice2), cap(newSlice2))
	fmt.Println(newSlice,iniArray,iniSlice)

	goods := [7]string{"sabun", "shampo", "sikat gigi", "pasta gigi", "sikat rambut", "sabun mandi"}
	goodsSlice := goods[1:4]
	fmt.Println(goodsSlice)
	newGoodsSlice := goods[0:2]
	appendNewGoodsSlice := append(newGoodsSlice, "sabunnn")
	fmt.Println(appendNewGoodsSlice)
	fmt.Println(goods)
	fmt.Println(goodsSlice)
}