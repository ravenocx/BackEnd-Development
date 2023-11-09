package main

import (
	"fmt"
	"time"
)

/*
Channels are pipes that link between goroutines
within your Go based applications that allow communication and
subsequently the passing of values to and from variables.
Jadi dengan channel dapat melakukan komunikasi antara goroutines
*/

func sendValue(c chan int){
	fmt.Println("Executing Goroutines")
	time.Sleep(1 * time.Second)
	c <- 8 //send operator -> 8 was sent to c
	fmt.Println("finished executing")
}

func OnlyIn(channel chan<- string){
	time.Sleep(2 * time.Second)
	channel <- "Haris"
}

func OnlyOut(channel <-chan string){
	data := <-channel
	fmt.Println(data)
}

func main() {
	fmt.Println("Channel")

	// values := make(chan int)
	/*
	With traditional unbuffered (undefined size chan) channels, whenever one goroutine sends a value to the channel, 
	that goroutine will subsequently block until the value is received from the channel.
	
	*/
	values := make(chan int,2) // buffered channel
	//These buffered channels are essentially queues of a given size that can be used for cross-goroutine communication
	// send operator should only block if that channel is full of capacity

	defer close(values) 
	//defer the closing of the channel to the end of the function
	// it stops or prevents any goroutines from sending or receiving to that channel one set is closed

	go sendValue(values)
	go sendValue(values)
	

	value := <- values // receive operator -> yang dikanan harus berupa channel
	fmt.Println(value)
	// using receive operator is to block until value is sent to the channel

	time.Sleep(1 * time.Second)

	/*
	Rangkuman
	1. Channel adalah tempat komunikasi untuk passing value sesama goroutine (synchronus)
	make(chan int)
	Data -> disimpan di Channel
	Channel -> diambil dari goroutine yang lain 
	
	Channel <- data (send operator)
	data := <- Channel (receive operator)

	2. Saat goroutine mengirim data ke channel, goroutine tsb akan terblock sampai ada yang menerima data tersebut
	jadi kalo ada goroutine yang memberi value ke channel harus di receive terlebih dahulu
	ketika ada 2 goroutine , maka send operator pada goroutine yang kedua akan terblock karena menunggu value untuk di received
	dan program akan otomatis berhenti sebelum send operator (goroutine tidak full tereksekusi)

	fmt.Println("Executing Goroutines")
	time.Sleep(1 * time.Second)
	c <- 8 --->> (akan stop sblm send operator)
	fmt.Println("finished executing")

	values := make(chan int,2) // buffered channel
	//These buffered channels are essentially queues of a given size that can be used for cross-goroutine communication
	// send operator should only block if that channel is full of capacity

	3. Ketika channel sudah tidak digunakan , harus di close() atau bisa menyebabkan memory leak
	it stops or prevents any goroutines from sending or receiving to that channel one set is closed

	*/

}