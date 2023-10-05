package main

import (
	"fmt"
	"errors"
)

func pembagian(penyebut int, pembagi int) (int,error){ // error -> interface (jadi bisa nil)
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}else {
		result := penyebut/pembagi
		return result, nil
	}
}

func main(){
	hasil,err := pembagian(100,10)
	
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}else {
		fmt.Println("Hasil", hasil)
	}
}