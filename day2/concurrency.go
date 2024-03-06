package day2

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second) // Introduce a delay of 1 second
	}
}

func GoRoutinesExample() {
	// Start a goroutine to execute the printNumbers function concurrently
	go printNumbers()

	// Continue executing other code in the main function
	for i := 1; i <= 3; i++ {
		fmt.Println("Main function:", i)
	}

	fmt.Println("Main function ended.")
}

func sender(ch chan<- int) {

	for i := 1; i <= 5; i++ {
		fmt.Println("Sending", i)
		// Send
		ch <- i // Send value into the channel
		time.Sleep(time.Second)
	}
	close(ch) // Close the channel when done sending
}

func receiver(ch <-chan int) {
	for {

		// Receive
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed. Receiver exiting.")
			break
		}
		fmt.Println("Received", val)
	}
}

func ChannelWithSenderAndReceiver() {
	// unbuffered channel
	ch := make(chan int)

	go sender(ch)

	go receiver(ch)

	// Wait for sender and receiver goroutines to finish
	// Better way is to use wait groups
	time.Sleep(7 * time.Second)
}

func ChannelWithDeadlock() {
	ch := make(chan int)

	// Blocked because there's no receiver
	ch <- 42

	// Blocked because there's no sender
	<-ch

	//TODO Resolve deadlock
}

func BufferedChannelExample() {
	ch := make(chan int, 3)

	// Send values into the buffered channel
	ch <- 1
	ch <- 2
	ch <- 3

	// what happens if send another value in the channel
	//ch <- 4

	// Receive values from the buffered channel
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
}

func ChannelSynchronization() {
	// Create a channel for synchronization
	done := make(chan int)

	// Goroutine 1: Sends value to channel
	go func() {
		defer close(done)
		fmt.Println("in sender")
		time.Sleep(5 * time.Second)
		done <- 1
	}()

	// Block the main goroutine until both goroutines finish
	fmt.Println("received ", <-done)

}

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
		close(ch) // Close the channel when all values are sent
	}()

	// Goroutine 2: Receiving values from the channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch { // Use range to iterate over channel values until it's closed
			fmt.Println("Received", val)
		}
	}()

	wg.Wait()
}

func ChannelsWithSelectStatement() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//  Sends int value to ch1 after 2 seconds
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 1
	}()

	// Goroutine 2: Sends int value 2 to ch2 after 3 seconds
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 2
	}()

	// Use select to wait for data on ch1 or ch2
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	case <-time.After(4 * time.Second):
		fmt.Println("Timed out waiting for data")
	}
}

func NonBlockingChannels() {

	ch := make(chan int)

	// Non-blocking receive
	select {
	case val := <-ch:
		fmt.Println("Received value from ch:", val)
	default:
		fmt.Println("ch is empty.")
	}

	// Non-blocking send (because there's space in the buffer)
	select {
	case ch <- 42:
		fmt.Println("Sent value to ch")
	default:
		fmt.Println("ch buffer is full")
	}

	// Non-blocking receive
	select {
	case val := <-ch:
		fmt.Println("Received value from ch:", val)
	default:
		fmt.Println("ch is empty")
	}
}

func CloseChannels() {
	done := make(chan bool)

	// Create an unbuffered channel
	ch := make(chan int)

	// Start a sender goroutine to send data to the channel
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()

	// Start a receiver goroutine to receive data from the channel
	go func() {
		defer func() {
			done <- true // Signal completion
			close(done)
		}()
		for {
			val, ok := <-ch
			if !ok {
				// If ok is false, the channel has been closed
				fmt.Println("Channel closed. Receiver exiting.")
				return
			}
			fmt.Println("Received:", val)
		}
	}()

	// Wait for both sender and receiver goroutines to finish

	fmt.Println("before done")
	<-done

	fmt.Println("after done")
}

func RangeOverChannels() {
	ch := make(chan int)

	// send routine
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}()

	// Iterate over channel
	for val := range ch {
		fmt.Println("Received:", val)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func WorkerPoolExample() {

	const numWorkers = 3
	const numJobs = 5

	// Create channels for sending jobs and receiving results
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(w)
	}

	// Send jobs to the workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	// Close results
	close(results)

	// Collect and print results
	for result := range results {
		fmt.Println("Result:", result)
	}

}

func waitGroupWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d: Starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d: Finished\n", id)
}

func WaitGroup() {
	var wg sync.WaitGroup // Initialize a WaitGroup

	numWorkers := 5
	wg.Add(numWorkers) // Add the number of goroutines to wait for

	// Start multiple goroutines
	for i := 1; i <= numWorkers; i++ {
		go waitGroupWorker(i, &wg)
	}

	fmt.Println("for completed")

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All workers have completed their tasks.")
}

func incrementCounter(counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(counter, 1)
	}
}

func AtomicInt() {
	var counter int64

	const numGoroutines = 10
	var wg sync.WaitGroup

	// Start multiple goroutines
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go incrementCounter(&counter, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final value of the counter
	fmt.Println("Final counter value:", counter)
}

func incrementCounterWithMutex(counter *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		//Acquire mutex
		mutex.Lock()
		*counter = *counter + 1
		// Release
		mutex.Unlock()
	}
}

func MutexWithRoutines() {

	var counter int

	const numGoroutines = 10
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go incrementCounterWithMutex(&counter, &wg, &mutex)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final value of the counter
	fmt.Println("Final counter value:", counter)
}

// Stateful goroutine that increments a shared counter
func incrementCounterStateful(id int, counter chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		// Receive the current count from the channel
		count := <-counter
		// Increment the count
		count++
		// Send the updated count back to the channel
		counter <- count
		fmt.Printf("Goroutine %d: Counter = %d\n", id, count)
		time.Sleep(time.Second) // Simulate some work
	}
}

func StatefulGoRoutine() {
	var wg sync.WaitGroup

	// Create a buffered channel with capacity 1 to store the count
	counter := make(chan int, 1)

	// Start with an initial count of 0
	counter <- 0

	// Start multiple goroutines that increment the counter
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go incrementCounterStateful(i, counter, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the counter channel
	close(counter)

	// Print the final value of the counter
	finalCount := <-counter
	fmt.Println("Final counter value:", finalCount)
}

func Signals() {
	// Create a channel to receive signals
	sigChan := make(chan os.Signal, 1)

	// Notify the sigChan channel on receiving specified signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-sigChan
	fmt.Printf("Received signal: %v\n", sig)

	// Perform cleanup or other tasks based on the received signal
	// For example, you can gracefully shut down the application here
	os.Exit(0)
}
