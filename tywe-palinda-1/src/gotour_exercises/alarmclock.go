package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	reminder := time.NewTicker(delay)
	for _ = range reminder.C {
		fmt.Println("The time is " + time.Now().Format("15:04:05") + ":  " + text)
	}
}

func main() {
	// go Remind("Time to eat", 10*time.Second)
	// go Remind("Time to work", 30*time.Second)
	// go Remind("Time to sleep", 60*time.Second)
	// select {}

	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
}
