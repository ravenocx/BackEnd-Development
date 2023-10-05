package main

import "fmt"

func random() interface{}{
	return "OK"
}

func main(){
	// type assertion -> kemampuan mengubah tipe data , biasanya digunakan terhadap interface{}
	result:= random()
	resultString := result.(string) //type assertion
	fmt.Println(resultString) 

	// resultInt := result.(int)//panic
	switch value:= result.(type){ // handle supaya tipe datanya sesuai
	case string :
		fmt.Println("String", value)
	case int :
		fmt.Println("Int", value)
	default :
		fmt.Println("Unknown", value)
	}

}