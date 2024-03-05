package day2

import (
	"fmt"
	"sync"
)

func Question2() {
	var (
		numStudents = 3
		sum         = 0
		wg          sync.WaitGroup
		mutex       sync.Mutex
	)

	for i := 1; i <= numStudents; i++ {
		wg.Add(1)
		go func(rating int, sum *int, mutex *sync.Mutex) {
			defer wg.Done()
			mutex.Lock()
			*sum += rating
			mutex.Unlock()

		}(i%6, &sum, &mutex)
	}
	wg.Wait()

	fmt.Println("Sum Rating:", sum)
	fmt.Println("Average Rating:", sum/numStudents)
}
