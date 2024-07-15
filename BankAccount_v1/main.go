package main

import (
	"fmt"
	"sync"
	"time"
)

var pf = fmt.Printf

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}
func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.balance < v {
		pf("Cannot withdraw: %d because balance is: %d\n", v, a.balance)
	} else {
		pf("Balance before: %d ", a.balance)
		a.balance -= v
		pf("Withdrawn: %d; Balance after: %d\n", v, a.balance)
	}
}

func main() {

	var acc Account
	acc.balance = 100

	for i := 0; i < 12; i++ {
		go acc.Withdraw(10)
	}
	time.Sleep(2 * time.Second)

}
