package main

import (
	"fmt"
	"strconv"
)

func main() {

	truth := true
	navigate(&truth)
	fmt.Println(truth)
	strconv.ParseInt

}

func navigate(value *bool) {

	*value = !*value

}
