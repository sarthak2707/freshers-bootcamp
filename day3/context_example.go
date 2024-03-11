package day3

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) {
	// Simulate a long-running operation
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Println("Operation canceled:", ctx.Err())
	}
}

func ContextExample() {
	// Create a context with a timeout of 1 second
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start a goroutine to perform the operation
	go process(ctx)

	// Wait for a few seconds to observe the operation
	time.Sleep(2 * time.Second)
}
