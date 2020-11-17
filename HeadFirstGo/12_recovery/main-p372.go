package main

import "fmt"

func calmDown() {
	p := recover()
	// fmt.Println(p.Error())

	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

func main() {
	defer calmDown()
	err := fmt.Errorf("there is an error")
	panic(err)
}
