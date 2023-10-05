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

	for { // endless loop
		select{ // which case triggers first is going to be executed
			case rec1 := <- one : //waiting for receiving one 
				fmt.Println("I received from channel one", rec1)
			case rec2 := <- two :
				fmt.Println("I received from channel two", rec2) // executed
			case <- quit : // if this case triggered the loop is going to stop
			// this case is wait for some channel received information from quit channel
				return
		}
	}
	

}