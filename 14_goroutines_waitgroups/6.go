package main

import (
	"fmt"
	"sync"
	// "time"
)

//goroutines
//launch function concurrently

func task(id int, w *sync.WaitGroup) {
	defer w.Done() //runs at last
	fmt.Println("doing task", id)
	//w.Done() another way
}

func main() {
	var wg sync.WaitGroup //waitgroup waits for a goroutine to finish the execution (solves no output without delay)
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg) //no output, since when main function ends, all goroutines are removed from scheduler
		// go func (i int){
		// 	fmt.Println(i)
		// }(i)
	}
	wg.Wait()
	//time.Sleep(time.Second*2) //delay to enable goroutines to run
}
