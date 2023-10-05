package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"` // json :"(key)"
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

type SensorReading struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Time        string `json:"time"`
	Information Info   `json:"info"`
}

type Info struct {
	Description string `json:"desc"`
}

func main() {
	book := Book{
		Title: "Learning go programming",
		Author: Author{
			Name:      "Haris",
			Age:       12,
			Developer: true,
		},
	}

	fmt.Printf("%+v\n", book)

	// byteArray, e := json.Marshal(book) // return byteArray with JSON format (encode) dengan bentuk yang agak berantakan

	byteArray, err := json.MarshalIndent(book, "", "  ") // best practice
	//json.MarshalIndent(struct, prefix, indent)
	/*
		prefix -> awalan baris
		indent -> berapa banyak space
	*/
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteArray)) // convert byte array

	// =============================================
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z", "info":{
		"desc" : "a sensor reading"
	}}`

	var reading SensorReading
	// var reading map[string]interface{}                 // jika tidak tau struktur field dari JSON yang ada
	er := json.Unmarshal([]byte(jsonString), &reading) // decode from json to struct
	// 1st parameter -> json and 2nd parameter -> struct targer
	if er != nil {
		panic(er)
	}
	fmt.Printf("%+v\n", reading)
}
