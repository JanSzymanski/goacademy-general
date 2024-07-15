package main

import (
	"fmt"
	"sync"
)

func main() {
	var value int
	dataStream := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for inc := range dataStream {
			fmt.Printf("Incrementing %d with %d\n", value, inc)
			value += inc
		}
		fmt.Println("Channel closed")
	}()
	go func(ch chan<- int) {
		defer wg.Done()
		defer close(dataStream)
		var wg2 sync.WaitGroup
		wg2.Add(2)

		go func() {
			defer wg2.Done()
			for i := 2; i <= 10; i += 2 {
				ch <- i
			}
		}()
		go func() {
			defer wg2.Done()
			for i := 1; i < 10; i += 2 {
				ch <- i
			}
		}()
		wg2.Wait()
	}(dataStream)
	wg.Wait()
}
