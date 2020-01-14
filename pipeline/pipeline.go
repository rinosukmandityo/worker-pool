package main

import (
	"fmt"
)

func main() {
	c := gen(1, 2, 3, 4, 5)
	for res := range seq(c) {
		fmt.Println(res)
	}

}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func seq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}
