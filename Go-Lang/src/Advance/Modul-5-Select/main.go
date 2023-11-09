package main

import (
	"fmt"
	"time"
)

func main() {
	// select -> allow us to wait for multiple communication operations of goroutines
	// Select statement waits until the communication(send or receive operation) is prepared for some cases to begin. 
	one := make(chan string)
	two := make(chan string)
	quit := make(chan string)

	go func(){
		time.Sleep(2 * time.Second)
		one <- "Hey"
	}()

	go func(){
		time.Sleep(1 * time.Second)
		two <- "Hello"
	}() // ()-> tanda bahwa functionnya langsung dipanggil ketika deklarasi

	// for -> if we want to wait for all of them
	for x:=0 ; x< 2 ; x++ {
		select{ // which case triggers first is going to be executed
			case rec1 := <- one : //waiting for receiving one 
				fmt.Println("I received from channel one", rec1)
			case rec2 := <- two :
				fmt.Println("I received from channel two", rec2) // executed
		}
	}
	
	counter := 0
	for { // endless loop
		select{ // which case triggers first is going to be executed
			case rec1 := <- one : //waiting for receiving one 
				fmt.Println("I received from channel one", rec1)
				counter++
			case rec2 := <- two :
				fmt.Println("I received from channel two", rec2) // executed
				counter++
			case <- quit : // if this case triggered the loop is going to stop
			// this case is wait for some channel received information from quit channel
				return
			default : // this case is triggered when there is no communication between goroutines
				fmt.Println("Menunggu data") 

			if counter == 2 {
				break
			}
		}
	}
	
	/*
	// select -> allow us to wait for multiple communication operations of goroutines
	// Select statement waits until the communication(send or receive operation) is prepared for some cases to begin. 

	Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, 
	jika data datang secara bersamaan di beberapa channel, maka akan dipilih secara random

	Menggunakan loop karena kita tidak tahu kapan antar goroutine berkomunikasi
	Jadi dengan for loop kita bisa looping sampai kondisi antar goroutine berkomunikasi
	*/

}