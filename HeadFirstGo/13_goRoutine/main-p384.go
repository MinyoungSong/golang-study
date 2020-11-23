package main

import (
	"fmt"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Print("b")
	}
}

func main() {

	go a()
	go b()
	fmt.Println("\nend main()")
}
