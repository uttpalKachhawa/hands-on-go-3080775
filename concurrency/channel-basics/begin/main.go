// concurrency/channels/begin/main.go
package main

import (
	"fmt"
	"time"
)

// sum calculates and prints the sum of numbers
func sum(nums []int, ch chan<- int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println("Result:", sum)
	// send the result to the channel
	ch <- sum
}


func main() {
	nums := []int{1, 2, 3, 4, 5}
	ch := make(chan int)
	// invoke the sum function as a goroutine
	go sum(nums, ch)

	result := <-ch // receive the result from the channel
	fmt.Println("Received from channel:", result)
 //create a buffered channel with capacity 1
	ch2 := make(chan string, 2) 
	ch2 <- "Hello, Channel!" 
	ch2 <- "How are you?"
	println(<-ch2) // Output: Hello, Channel!
	println(<-ch2) // Output: How are you?
	// send a message to the buffered channel
	// force main thread to sleep
	time.Sleep(100 * time.Millisecond)
}
