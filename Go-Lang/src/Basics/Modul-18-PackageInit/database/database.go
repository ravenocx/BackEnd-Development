package database

import "fmt"

var connection string

func init() {
	// init -> func yang dieksekusi pertama kali dan akan selalu dieksekusi ketika package diakses
	fmt.Println("Init dipanggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}