package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func main() {
	objPerson := []Person{}

	fmt.Println(len(objPerson))
}