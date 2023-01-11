package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

/*
Balance() -> N at a time
Deposit() -> Just once at a time
*/

/*
Write Lock or Regular Lock
*/
func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Withdraw(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b - amount
	lock.Unlock()
}

/*
Read lock
*/

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go Withdraw(i*50, &wg, &lock)
	}

	wg.Wait()
	fmt.Println(Balance(&lock))
}
