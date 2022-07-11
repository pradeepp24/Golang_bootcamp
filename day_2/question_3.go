package main

import (
	"fmt"
	"sync"
)

type bankAccount struct {
	m       sync.Mutex
	balance int
}

func (acc *bankAccount) deposit(x int) {
	acc.m.Lock()
	acc.balance += x
	defer acc.m.Unlock()
}

func (acc *bankAccount) withdraw(x int) {
	acc.m.Lock()
	defer acc.m.Unlock()
	if x > acc.balance {
		fmt.Println("insufficient balance")
	} else {
		acc.balance -= x
	}
}

func main() {
	account := bankAccount{balance: 500}
	var wg sync.WaitGroup

	doOperation := func(operation string, amt int) {
		if operation == "withdraw" {
			account.withdraw(amt)
		} else if operation == "deposit" {
			account.deposit(amt)
		} else {
			fmt.Println("invalid operation")
		}
		wg.Done()
	}

	wg.Add(5)
	go doOperation("deposit", 1000)
	go doOperation("withdraw", 1000)
	go doOperation("withdraw", 1000)
	go doOperation("withdraw", 1000)
	go doOperation("block", 1000)
	wg.Wait()
	fmt.Println(account.balance)
}
