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
	time.Sleep(100 * time.Millisecond)
}

func concurrencyWithChannels() {
	channel := make(chan bool)
	go printSlowWithChannel(channel)
	fmt.Println("I am ahead of print slow")
	<-channel

}

func bufferingConncurrency() {
	// messages := make(chan string)
	messages := make(chan string, 1)

	messages <- "Hello world"
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
