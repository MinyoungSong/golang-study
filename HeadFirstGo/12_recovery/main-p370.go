package main

import "fmt"

func calmDown() {
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
