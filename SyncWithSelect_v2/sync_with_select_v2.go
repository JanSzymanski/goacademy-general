package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	ch1 := make(chan int)
	ch2 := make(chan int)
	var counter int
	maxIter := 10

	go func() {
		defer wg.Done()
		for i := 2; i <= maxIter; i += 2 {
			ch1 <- i
		}
		ch1 = nil
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < maxIter; i += 2 {
			ch2 <- i
		}
		ch2 = nil
	}()
	go func() {
		defer wg.Done()
		for range maxIter {
			select {
			case val := <-ch1:
				fmt.Println("Updating counter: ", counter, " with: ", val)
				counter += val
			case val := <-ch2:
				fmt.Println("Updating counter: ", counter, " with: ", val)
				counter += val
			}
		}
	}()
	wg.Wait()
	fmt.Println("Program finished")
}
