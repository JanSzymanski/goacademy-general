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

	go func() {
		defer wg.Done()
		defer close(ch1)
		for i := 2; i <= 10; i += 2 {
			ch1 <- i
		}
	}()
	go func() {
		defer wg.Done()
		defer close(ch2)
		for i := 1; i < 10; i += 2 {
			ch2 <- i
		}
	}()
	go func() {
		defer wg.Done()
		for ch1 != nil || ch2 != nil {
			select {
			case val, open := <-ch1:
				if open {
					fmt.Println("Updating counter: ", counter, " with: ", val)
					counter += val
				} else {
					fmt.Println("Channel 1 closed, putting nil")
					ch1 = nil
				}

			case val, open := <-ch2:
				if open {
					fmt.Println("Updating counter: ", counter, " with: ", val)
					counter += val
				} else {
					fmt.Println("Channel 2 closed, putting nil")
					ch2 = nil
				}
			}
		}
	}()
	wg.Wait()
	fmt.Println("Program finished")
}
