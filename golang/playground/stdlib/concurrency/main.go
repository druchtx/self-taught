package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	fmt.Printf("start main: %s \n", time.Now().Format(time.DateTime))

	ch := make(chan struct{})
	done := make(chan struct{})
	sema := make(chan struct{}, 5)
	var wg sync.WaitGroup

	// Start the producer
	go assign(ch, done)

	// Start the signal listener. It will close the 'done' channel.
	go func() {
		terminate := make(chan os.Signal, 1)
		signal.Notify(terminate, os.Interrupt)
		<-terminate // Block until signal
		fmt.Printf("main: received interrupt, starting graceful shutdown...\n")
		close(done)
	}()

	// The consumer loop. It ranges over 'ch' until it's closed by 'assign'.
	for range ch {
		// For each task, add to the WaitGroup and acquire the semaphore.
		wg.Add(1)
		sema <- struct{}{}
		go func() {
			// Defer releasing the semaphore and marking the WaitGroup as done.
			defer func() {
				<-sema
				wg.Done()
			}()
			fmt.Printf("start work: %s \n", time.Now().Format(time.DateTime))
			work()
		}()
	}

	// After 'ch' is closed, the loop above exits.
	// Now, wait for all running worker goroutines to finish.
	fmt.Println("main: task channel closed, waiting for workers to finish..")
	wg.Wait()

	fmt.Printf("stop main: graceful shutdown complete: %s \n", time.Now().Format(time.DateTime))
}

// work simulates a long running task.
func work() {
	time.Sleep(time.Second * 20) // Let's use a shorter time for demonstration
	fmt.Printf("finish work: %s \n", time.Now().Format(time.DateTime))
}

// assign is the producer. It sends tasks every second until the 'done' channel is closed.
func assign(ch chan<- struct{}, done <-chan struct{}) {
	// Ensure the channel is closed when this function exits.
	defer func() {
		fmt.Println("assign: closing task channel.")
		close(ch)
	}()

	for {
		select {
		case <-time.After(time.Second * 1):
			// Non-blocking send to avoid getting stuck if the main consumer isn't ready
			// This can happen during shutdown.
			select {
			case ch <- struct{}{}:
				fmt.Printf("assign work: %s \n", time.Now().Format(time.DateTime))
			case <-done:
				fmt.Println("assign: stopping production during send.")
				return
			}
		case <-done:
			fmt.Println("assign: stopping production.")
			return
		}
	}
}
