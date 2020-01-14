package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// CancelString()
	CancelRandom()
}

func CancelRandom() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()

		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)

	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)
	time.Sleep(time.Second * 1)

}

func CancelString() {
	doWork := func(done <-chan interface{}, data <-chan string) <-chan interface{} {
		terminated := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)

			for {
				select {
				case <-done:
					return
				case s := <-data:
					fmt.Println(s)
				}
			}
		}()

		return terminated
	}

	done := make(chan interface{})
	inputData := make(chan string)
	terminated := doWork(done, inputData)

	go func() {
		inputData <- "inputData"
		time.Sleep(time.Second * 1)
		fmt.Println("Canceling doWork goroutine from parent...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}
