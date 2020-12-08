package main

import (
	"automated/testing/prose"
	"fmt"
)

func main() {

	input := []string{"apple", "orange", "grape"}
	result := prose.JoinWithCommas(input)

	fmt.Println(result)
}
