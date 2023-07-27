package main

import (
	"fmt"
	"time"
)

func main() {

	// basicConcurrency()
	// concurrencyWithChannels()
	// bufferingConncurrency()
	loopingOverChannels()

}

func loopingOverChannels() {
	someNums := make(chan int)
	go fillNums(someNums)
	for num := range someNums {
		fmt.Println(num)
	}

}

func basicConcurrency() {
	go printSlow()
	time.Sleep(20 * time.Millisecond)
	fmt.Println("I am ahead of print slow")
	// time.Sleep(100 * time.Millisecond)
}

func concurrencyWithChannels() {
	channel := make(chan bool)
	go printSlowWithChannel(channel)
	fmt.Println("I am ahead of print slow")
	<-channel

}

func bufferingConncurrency() {
	// If given no arguments the channel is unbuffered
	// This means that adding a value to a channel is blocking and will only be unblocked when another goroutine removes that value
	messages := make(chan string)

	// The second argument to the make command for a channel is its buffering capacity
	// messages := make(chan string, 1)

	// This will cause the program to crash if the channel is not given a buffering capacity
	messages <- "Hello world"

	// A second value added into the channel will cause the program to crash if the buffering capacity is only 1
	// messages <- "Hello world"
	fmt.Println(<-messages)

}

func printSlow() {

	fmt.Println("Print slow started")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("I print slow")
}

func printSlowWithChannel(channel chan bool) {
	printSlow()
	channel <- true

}

func fillNums(channel chan<- int) {

	for i := 1; i < 11; i++ {
		channel <- i
	}

	close(channel)

}
