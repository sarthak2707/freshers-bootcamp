package day2Exercises

import (
	"fmt"
	"sync"
)

type bankAccount struct {
	accountNumber string
	balance       float32
}

func (ba *bankAccount) credit(amount float32, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	ba.balance += amount
	fmt.Printf("amount %f credited to account - %s\n", amount, ba.accountNumber)
}

func (ba *bankAccount) debit(amount float32, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	if newBalance := ba.balance - amount; newBalance < 0 {
		fmt.Println("can't withdraw more money than the current balance")
		return
	}
	ba.balance -= amount
	fmt.Printf("amount %f debited from account - %s\n", amount, ba.accountNumber)
}

func BankBalanceUpdate() {
	var ba bankAccount
	var mutex sync.Mutex
	var wg sync.WaitGroup

	ba.accountNumber = "89417058301020"
	ba.balance = 500

	wg.Add(3)
	go ba.credit(100, &mutex, &wg)
	go ba.debit(1000, &mutex, &wg)
	go ba.debit(300, &mutex, &wg)

	wg.Wait()
	fmt.Printf("final balance - %f\n", ba.balance)
}
