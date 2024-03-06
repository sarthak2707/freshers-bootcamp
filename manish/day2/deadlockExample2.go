package day2

import (
	"fmt"
	"sync"
)

func ChannelWithDeadlockExample2() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Goroutine 1: Sending values into the channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			fmt.Println("Sending", i)
			ch <- i // Send value into the channel
		}
		close(ch)
	}()

	// Goroutine 2: Receiving values from the channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			val, ok := <-ch // Receive value from the channel
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Received", val)
		}
	}()

	wg.Wait()

	// TODO Resolve Deadlock
}
