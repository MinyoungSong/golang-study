package stack

import "fmt"

func stack(array []string) {

	stackArr := []string{}

	for _, v := range array {
		stackArr = append(stackArr, v)
		fmt.Printf("stackArr :: %v\n", stackArr)
	}

	for i := 0; i < len(array); i++ {

		pop := stackArr[len(stackArr)-1]
		stackArr = stackArr[:len(stackArr)-1]

		fmt.Printf("pop :: %v\n", pop)
		fmt.Printf("stackArr :: %v\n", stackArr)
	}

}
