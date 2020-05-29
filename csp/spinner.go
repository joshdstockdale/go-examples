package main

import (
	"fmt"
	"time"
)

func main(){
	// go spinner runs concurrently
	go spinner(100*time.Millisecond)
	const n = 45
	// uses inefficient recursive algorithm
	fibN := fib(n)
	fmt.Printf("\rFibonacci (%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration){
	for {
		for _, r := range `-\|/`{
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int{
	if x < 2{
		return x
	}
	return fib(x-1) + fib(x-2)
}
