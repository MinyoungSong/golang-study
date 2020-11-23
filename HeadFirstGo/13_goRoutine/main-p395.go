package main

import (
	"fmt"
	"time"
)

func main() {
	channelA := make(chan int)
	channelB := make(chan int)

	go odd(channelA)
	go even(channelB)

	fmt.Println(<-channelA)
	fmt.Println(<-channelA)
	fmt.Println(<-channelB)
	fmt.Println(<-channelB)
	time.Sleep(3 * time.Second)

}

func odd(channel chan int) {
	channel <- 1
	fmt.Println("send 1")
	channel <- 3
	fmt.Println("send 3")
}

func even(channel chan int) {
	channel <- 2
	fmt.Println("send 2")
	channel <- 4
	fmt.Println("send 4")
}
