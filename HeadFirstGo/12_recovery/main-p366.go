package main

import (
	"fmt"
)

func one() {
	defer fmt.Println("defered in one()")
	two()
}

func two() {
	defer fmt.Println("defered in two()")
	panic("Let's see what's been deferred")
}

func main() {
	one()
}
