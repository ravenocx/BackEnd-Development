package main

import "fmt"

type Address struct {
	City,Provincy,Country string
}

func main(){
	address1 := Address{"Subang","Jawa Barat","Indonesia"}
	address2 := address1 // Pass by value 
	// Pass by value -> ketika mengirimkan variable ke tempat yg lain, sebenarnya yg dikirim adalah duplikasinya di memori
	// go secara default pass by value bukan pass by reference

	address2.City = "Bandung"

	fmt.Println(address1) // address 1 tidak berubah
	fmt.Println(address2)

	// Pointer -> kemampuan membuat reference ke lokasi data di "memori yang sama" tanpa menduplikasi data
	// atau yg dikenal sebagai pass by reference

	var address3 *Address = &address1 // pointer ke address 1 , mereference ke address 1
	fmt.Println(address3)

	address3.City = "Bekasi"
	address3.Provincy = "Planet namex"
	address3.Country = "Nigeria"
	
	fmt.Println(address1)

	address3 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} //tidak berpengaruh ke address 1 
	// Address 1 tidak berubah karena membuat variable baru yang mereferensi ke variabel yang berbeda
	// dia mereferensi ke memori yang baru -> &Address{"Jakarta", "DKI Jakarta", "Indonesia"} , dengan tempat memori yg berbeda dgn address 1


	// Saat mengubah variable pointer (secara keseluruhan), maka yang berubah hanya variable tersebut
	// semua variable yang mengacu ke data yang sama tidak berubah

	fmt.Println(address1)
	fmt.Println(address3)

	// & -> buat pointer variable, atau buat menginialisasi data variabel
	// * -> pointer ke tipe data (buat deklarasi)

	address4 := &address1
	address5 := &address1
	fmt.Println(address1)
	fmt.Println(address5)
	*address4 = Address{"Malang", "Jawa Timur", "Indonesia"} // add1 dan add5 ikut berubah referencenya(pointer)
	fmt.Println(address1) // kepala pointernya tetap di address 1
	fmt.Println(address4)
	fmt.Println(address5)

	alamat := new(Address) // membuat pointer dengan data kosong -> disebut nil pointer
	fmt.Println(alamat)

	/*
	Kesimpulan : 
	Kegunaan operator asterisk (*) di dalam hal ada tiga:
		1. Untuk mendeklarasikan variabel pointer.
		Contoh: var address2 *Address = &address1

		2. Untuk mengambil value dari variabel pointer. Contohnya jika kita melakukan printout ke layar 
		tidak ada tanda ampersand (&).
		Contoh: fmt.Println(*address2)

		3. Untuk memindahkan referensi dari kedua variabel (variabel pertama dan variabel kedua yang 
		berupa pointer yang menunjuk ke alamat variabel pertama) kepada alamat referensi dari nilai yang baru.
		Contoh: *address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

		The & operator is used to find the address of a variable. The * operator is used to access the 
		value at the address specified by a pointer. The * operator is also used to declare a pointer variable.

		& = referencing = mengubah value ke memory
		* = dereferencing = mengubah memory ke value
	*/

}