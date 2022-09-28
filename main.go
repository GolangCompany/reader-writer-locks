package main

import (
	"fmt"
	"sync"
)

var (
	mutex  sync.Mutex
	amount int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depositing %d to the account with balance %d\n", value, amount)
	amount += value
	mutex.Unlock()
	wg.Done()
}
func extract(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance %d\n", value, amount)
	amount -= value
	mutex.Unlock()
	wg.Done()
}
func main() {
	fmt.Println("Hello Everyone")
	amount = 1000
	var wg sync.WaitGroup
	wg.Add(2)
	go extract(600, &wg)
	go deposit(500, &wg)
	wg.Wait()
	fmt.Printf("New Balance %d\n", amount)
}
