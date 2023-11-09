package main

import (
	"fmt"
	"sync"
)

/*
Go allows us to run code concurrently using goroutines. However, when concurrent processes access the same piece of data,
it can lead to race conditions. Mutexes are data structures provided by the sync package.
They can help us place a lock on different sections of data so that only one goroutine can access it at a time.
With mutex -> we can write goroutine in safe

dimana ketika kita melakukan locking terhadap mutex, 
maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan unlock

variable yang di lock access nya adalah diantara method lock() dan unlock()

*/

var (
	mutex sync.Mutex
	balance int
)

func deposit(value int,wg *sync.WaitGroup){
	mutex.Lock() // lock the mutex -> the calling goroutine blocks until the mutex is available(unlock)
	fmt.Printf("Depositing %d to account with balance %d\n",value, balance)
	balance += value
	mutex.Unlock() 
	wg.Done() // telling the goroutine has finished the execution
}

func withdraw(value int, wg *sync.WaitGroup){
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}


func main(){
	balance = 1000
	var wg sync.WaitGroup // A WaitGroup waits for a collection of goroutines to finish
	wg.Add(2) //The main goroutine calls Add to set the number of goroutines to wait for. (berapa banyak goroutine yang ditunggu)
	// Terdapat wg karena ada yang namanya Deadlock (dimana antar goroutine menunggu lock sehingga tidak ada yang jalan)
	go withdraw(700, &wg)
	go deposit(500,&wg)
	wg.Wait() // block main function until all goroutine completed

	fmt.Printf("New balance %d",balance)
}