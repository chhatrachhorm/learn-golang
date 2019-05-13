package main

/*
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.
By default sends and receives block until both the sender and receiver are ready.
This property allowed us to wait at the end of our program for the transmission without having to use any other synchronization.
*/
import "fmt"

func foo(c chan int, value int)  {
	c <- value
}

func main()  {
	intChan := make(chan int)
	go foo(intChan, 152)
	fmt.Println("The channel item value is", <-intChan)
}
