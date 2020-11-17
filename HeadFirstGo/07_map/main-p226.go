package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello")

	// var medalMap map[string]int
	medalMap := make(map[string]int)
	medalMap["bronze"] = 3
	medalMap["silver"] = 2
	medalMap["gold"] = 1

	value, ok := medalMap["gold"]
	fmt.Println(value)
	fmt.Println(ok)

	value2, ok := medalMap["red"]
	fmt.Println(value2)
	fmt.Println(ok)

	fmt.Println(medalMap["blue"])

	for k, v := range medalMap {
		fmt.Printf("The %s medal's rank is %d \n", k, v)
	}

}
