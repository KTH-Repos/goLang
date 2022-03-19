package main

import (
	"fmt"
)

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int) {
	// TODO sum a
	// TODO send result on res
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	res <- sum
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	if len(a) != 0 {
		n := len(a)
		ch := make(chan int)
		go sum(a[:n/2], ch)
		go sum(a[n/2:], ch)
		// TODO Get the subtotals from the channel and return their sum
		sum1, sum2 := <-ch, <-ch
		return sum1 + sum2
	}
	return 0
}

func main() {
	// example call
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ConcurrentSum(a))
}
