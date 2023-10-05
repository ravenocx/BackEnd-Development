package entity

// mock adalah object tiruan untuk menggantikan dependensi/service 3rd party
// contoh : toko online, terdapat API call ke payment gateway
// kita tidak bisa melakukan test ke payment gateway (karena bukan sistem kita )

// unit test -> idealnya tidak merunning 3rd party system apapun

type Category struct{ // representasi database
	Id string
	Name string
}