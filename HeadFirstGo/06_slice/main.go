package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello")

	var testSlice []string
	testSlice = make([]string, 3)
	testSlice[1] = "two"

	fmt.Println(testSlice)
	fmt.Println(len(testSlice))

	notes := []string{"a", "b", "c"}
	fmt.Println(notes)

	notes = append(notes, "d")
	fmt.Println(notes)

	fmt.Println(os.Args)
	arguments := os.Args[1:]
	fmt.Println(arguments)

	serveralInts(1, 2, 3, 4, 5)

	result := average(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}

func serveralInts(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(numbers[1:])
}

func average(numbers ...float64) float64 {

	var total float64 = 0
	for _, number := range numbers {

		total += number
	}

	return total / float64(len(numbers))
}
