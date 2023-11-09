package main 

import "fmt"

func main(){
	identity := map[string]string{ // map[key]value
		"firstName" : "Muhammad",
		"lastName"	: "Haris",
		"Age"		: "19",
	}

	identity["city"] = "Jakarta"
	identity["Age"] = "120"
	fmt.Println(identity)
	fmt.Println(identity["Age"])

	// person := make(map[string]string) -> create new map
	// delete(map,key) -> delete data di map pada key yang ada
}