package main

import (
	"context"
	"fmt"
	"time"
)

/*
Konsep context memungkinkan Anda untuk membatalkan,
mengatur batas waktu, atau berbagi data dan informasi penting
antar goroutine dengan cara yang aman dan terstruktur.
Context biasanya dibuat per request
*/
func enrichContext (ctx context.Context) context.Context{
	// whenever you work with context, the context must be on 1st in parameter
	return context.WithValue(ctx, "request-id","12345") // with value -> membuat copy dari parent context/ original context
	// return copy dari parent context dengan key(request-id) value (12345) yang baru 
}

func doSomethingCool (ctx context.Context){
	rID := ctx.Value("request-id") // mengambil value dari request-id berdasarkan key yang ada di argument
	// digunakan untuk pass information within your application -> mirip seperti bucket
	// context hanya digunakan ketika informasi yang ingin disebarkan digunakan pada keseluruhan program
	// jangan menaruh semua informasi di context
	fmt.Println(rID)
	for{ // infinite loop
		select{
			case <- ctx.Done(): // waiting for a read from ctx.Done()
			// ctx.Done() -> for cancelation
				fmt.Println("timed out")
				return // to ensure cancel this function calling
			default :
				fmt.Println("doing something cool")
				// jika ctx.Done() belum di closed berdasarkan waktu deadline case default akan di eksekusi terus
		}
		time.Sleep(500 * time.Millisecond) // supaya terminalnya tidak terjadi spam
	}
	
}

func main(){
	// ctx := context.Background() //create new context -> original context
	ctx , cancel := context.WithTimeout(context.Background(),2 *time.Second) // creating new context with deadline
	/* ketika melewati time deadline -> the done channel within the context object is going to be closed (ctx.Done())
	dan membatalkan eksekusi doSomethingCool()
	*/
	defer cancel()
	// cancel() (tanpa defer didepannya) ->  if we want to cancel our app sooner than the timeout
	ctx = enrichContext(ctx)
	go doSomethingCool(ctx) // go routines -> asynchronous

	select{
	case <- ctx.Done(): // waiting for done channel to be closed
		fmt.Println("oh no, i've exceeded the deadline")
		fmt.Println(ctx.Err())
		// if we call Err() function and the Done channel is not yet closed -> we dont receive any error
		// if we call Err() function and the Done channel is closed -> return an error (the context has been cancelled)
	}

	time.Sleep(2 * time.Second)
	
	
}