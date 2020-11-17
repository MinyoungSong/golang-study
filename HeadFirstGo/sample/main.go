package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter a grade :")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Your grade is :", input)

	// var originalCount int
	// originalCount = 10

	// eatenCount := 4

	// fmt.Println("I started with", originalCount, "apples")
	// fmt.Println("Some jerk ate", eatenCount, "apples")
	// fmt.Println("There are", originalCount-eatenCount, "apples left")

}
