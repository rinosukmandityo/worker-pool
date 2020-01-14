package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hit int
}

func main() {
	table := make(chan *Ball)
	done := make(chan struct{})
	go Player("ping", table, done)
	go Player("pong", table, done)

	// table <- new(Ball) // game on; toss the ball
	time.Sleep(time.Second)
	done <- struct{}{} // game over; grab the ball
}

func Player(name string, table chan *Ball, done <-chan struct{}) {
	for {
		select {
		case ball := <-table:
			ball.hit++
			table <- ball
			fmt.Println(name, ball.hit)
			time.Sleep(100 * time.Millisecond)
		case <-done:
			fmt.Println("timeout reached")
		}
	}
}
