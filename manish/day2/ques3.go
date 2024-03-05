package day2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func depositMoney(balance *int, depositArray []int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, deposit := range depositArray {
		mutex.Lock()
		*balance += deposit
		fmt.Println("Balance after depositing ", deposit, " :", *balance)
		mutex.Unlock()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func withdrawMoney(balance *int, withdrawArray []int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, withdraw := range withdrawArray {
		mutex.Lock()
		if *balance < withdraw {
			fmt.Println("Insufficient Balance. Cannot withdraw ", withdraw, " from ", *balance)
			mutex.Unlock()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
			continue
		}
		*balance -= withdraw
		fmt.Println("Balance after withdraw ", withdraw, " :", *balance)
		mutex.Unlock()
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}
func Question3() {
	var balance = 500

	var depositArray = []int{200, 100, 300, 50, 100}
	var withdrawArray = []int{100, 200, 50, 200, 1000}

	var mutex sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go depositMoney(&balance, depositArray, &mutex, &wg)

	go withdrawMoney(&balance, withdrawArray, &mutex, &wg)
	wg.Wait()

	fmt.Println("Final Balance:", balance)
}
