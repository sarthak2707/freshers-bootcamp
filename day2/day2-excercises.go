package day2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// task 1
func countFrequency() {
	input := []string{"quick", "brown", "fox", "lazy", "dog"}

	result := make(map[rune]int)
	var wg sync.WaitGroup
	letterCh := make(chan rune)

	// Start goroutine for each string to count letters
	for _, str := range input {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			for _, char := range s {
				letterCh <- char
			}
		}(str)
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(letterCh)
	}()

	// Collect letter counts from channel
	for char := range letterCh {
		result[char]++
	}

	// Print the result
	fmt.Println("Output:")
	fmt.Println(result)
}

// task 2

func AverageRating() {
	rand.Seed(time.Now().UnixNano())

	numStudents := 200
	var wg sync.WaitGroup
	ratings := make(chan int, numStudents)

	// Simulate students providing ratings
	for i := 1; i <= numStudents; i++ {
		wg.Add(1)
		go func(studentID int) {
			defer wg.Done()
			rating := simulateRating()
			fmt.Printf("Student %d rated the teacher with %d\n", studentID, rating)
			ratings <- rating
		}(i)
	}

	// Close ratings channel once all ratings are received
	go func() {
		wg.Wait()
		close(ratings)
	}()

	// Calculate average rating
	totalRatings := 0
	numRatings := 0
	for rating := range ratings {
		totalRatings += rating
		numRatings++
	}

	// Print average rating
	averageRating := float64(totalRatings) / float64(numRatings)
	fmt.Printf("Average rating of the teacher: %.2f\n", averageRating)
}

// Simulate a rating between 1 and 10
func simulateRating() int {
	return rand.Intn(10) + 1
}

// task 3

type BankAccount2 struct {
	balance int
	mutex   sync.Mutex
}

func NewBankAccount(initialBalance int) *BankAccount2 {
	return &BankAccount2{
		balance: initialBalance,
	}
}

func (acc *BankAccount2) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	acc.balance += amount
	fmt.Printf("Deposited: Rs.%d\n", amount)
}

func (acc *BankAccount2) Withdraw(amount int, wg *sync.WaitGroup) bool {
	defer wg.Done()
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	if acc.balance-amount < 0 {
		fmt.Println("Withdrawal failed: Insufficient balance")
		return false
	}
	acc.balance -= amount
	fmt.Printf("Withdrawn: Rs.%d\n", amount)
	return true
}

func BankAccountOperations() {
	// Create a bank account with an initial balance of Rs. 500
	account := NewBankAccount(500)

	// Number of deposit and withdrawal operations to perform concurrently
	numOperations := 10

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Perform concurrent deposit and withdrawal operations
	for i := 0; i < numOperations; i++ {
		wg.Add(2)
		go account.Deposit(100, &wg)
		go account.Withdraw(50, &wg)
	}

	// Wait for all operations to finish
	wg.Wait()

	// Print the final balance
	fmt.Printf("Final Balance: Rs.%d\n", account.balance)
}
