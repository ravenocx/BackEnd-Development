package main

import (
	"fmt"
)

func newMap (name string) map[string]string{
	if name == "" {
		return nil
	}else {
		return map[string]string {
			"Name" : name,
		}
	}
}

func main(){
	// di golang ketika mendeklarasikan variable otomatis akan terbentuk default valuenya (bukan nil)
	/*
	default value
	int = 0
	bool = false
	string = ""
	struct -> isi fieldnya kosong
	*/
	// nil sendiri hanya bisa digunakan pada interface,function,map,slice,pointer dan channel
	data := newMap("")
	if data == nil {
		fmt.Println(data) // return nil (map[])
		fmt.Println("Data Kosong")
	}else {
		fmt.Println(data)
	}
	
}