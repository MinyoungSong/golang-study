package main

import "fmt"

func calmDown() {
	// 예외처리.......
	recover()
}

func freakOut() {
	// defer recover()
	defer calmDown()
	panic("Oh no")
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
