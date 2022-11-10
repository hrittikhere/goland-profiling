package main

import (
	"fmt"
	"sync"
	"time"
)

func Mutx(s int) int {
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	wg.Add(1000 * 1000)
	for i := 0; i < 1000*1000; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			s++
		}(i)
	}
	wg.Wait()
	return s
}

var ch = make(chan int)

func Block() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go G1(&wg)
	go G2(&wg)
	wg.Wait()
}

func G1(wg *sync.WaitGroup) {
	ch <- 100
	wg.Done()
}

func G2(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	<-ch
	wg.Done()
}

func FibRecursive(n int) uint64 {
	// Function Creation to accept integer till which the Fibonacci series runs
	if n == 0 {
		return 0
		// base case for the recursive call
	} else if n == 1 {
		return 1
		// base case for the recursive call
	} else {
		return FibRecursive(n-1) + FibRecursive(n-2)
		// recursive call for finding the fibonacci number
	}
}

var f [100010]uint64

func Fib(n int) uint64 {
	f[0] = 0
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}

func main() {
	fmt.Println(Fib(3))
}
