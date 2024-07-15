package main

import (
	"fmt"
	"sync"
	"time"
)

type Destination struct {
	lock  sync.RWMutex
	value int
}

func (d *Destination) Increment(inc int, act string) {
	d.lock.Lock()
	fmt.Printf("%s:\ti:%d\tUpdating value: %d -> %d\n", act, inc, d.value, d.value+inc)
	d.value += inc
	d.lock.Unlock()

}

func main() {
	dest := new(Destination)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(d *Destination) {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			d.Increment(i, "Even")
			time.Sleep(1 * time.Second)
		}
	}(dest)
	go func(d *Destination) {
		defer wg.Done()
		for i := 1; i < 10; i += 2 {
			d.Increment(i, "Odds")
			time.Sleep(1 * time.Second)
		}
	}(dest)
	wg.Wait()

}
