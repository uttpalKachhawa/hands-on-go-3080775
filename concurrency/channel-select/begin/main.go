// concurrency/channel-select/begin/main.go
package main

import "time"


type signal struct{}

func main() {
	// declare an empty struct channel for signaling

done := make(chan signal)
	// declare a timer channel
timer := time.NewTimer(1 * time.Second)
	// launch a goroutine to signal after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		done <- signal{} // send an empty struct to signal completion
	}()
	// wait for a signal on either channel
	select {
	case <-done:
		println("Received signal from done channel")
	case <-timer.C:
		println("Timer expired")
	}
}
