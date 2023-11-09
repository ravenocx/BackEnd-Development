package Once

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			once.Do(OnlyOnce)
			// once memastikan sebuah function di eksekusi hanya sekali
			//Jadi berapa banyak pun goroutine yang mengakses, bisa dipastikan 
			// bahwa goroutine yang pertama yang bisa mengeksekusi function nya
			// Goroutine yang lain akan di hiraukan, artinya function tidak akan dieksekusi lagi

			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}