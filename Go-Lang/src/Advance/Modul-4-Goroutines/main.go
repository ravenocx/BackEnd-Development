package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second) // before printing i dikasih jeda selama 1 detik
		fmt.Println(i)
	}
}
	/*
		Using goroutines is a very quick way to turn what would be a
		sequential program into a concurrent program without having to
		worry about things like creating threads or thread-pools.

		Goroutines are incredibly lightweight threads managed by the go runtime. They enable us to create
		asynchronous parallel programs that can execute some tasks far quicker than if they were written in a sequential manner.
	*/

func main() {
	fmt.Println("Goroutines")
	// compute(5)
	// compute(5) 
	// dua pemanggilan function diatas di run secara synchronous

	go compute(5) // -> goroutines , jadi eksekusi functionnya menjadi lebih cepat dan dieksekusi secara asynchronous
	go compute(5)

	// if we run the program, it will completes without printing the loop
	// this because main function complete before async func calls could execute
	// any goroutines that have yet to complete are promptly terminated.
	fmt.Scanln()// agar goroutines bisa dieksekusi
	
}