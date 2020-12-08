package main

import (
	"fmt"
	"strconv"
)

// Student is
type Student struct {
	ID      int
	Name    string
	Address string
	Age     int
}

func main() {

	var studentIDs = [5]int{1, 2, 3, 4, 5}
	var resultStudents []Student
	channel := make(chan Student) // bufferd channel 사용 시 지정한 buffer가 다 채워질때까지 non-block으로 실행됨.

	for _, id := range studentIDs {
		go executeTask(id, channel)
	}

	for i := range channel {

		fmt.Println(i)
		resultStudents = append(resultStudents, i)

		if len(resultStudents) == 5 {
			break
		}
	}

	// time.Sleep(2 * time.Second)
	fmt.Println(resultStudents)

}

func executeTask(id int, channel chan Student) {

	fmt.Println("executeTask " + strconv.Itoa(id))

	student := Student{
		ID:      id,
		Name:    "",
		Address: "",
		Age:     0,
	}

	subTaskA(&student)
	subTaskB(&student)
	subTaskC(&student)

	channel <- student
}

func subTaskA(student *Student) {
	// fmt.Println("subTaskA " + strconv.Itoa(student.ID))
	student.Name = "Andy" + strconv.Itoa(student.ID)

}

func subTaskB(student *Student) {
	// fmt.Println("subTaskB " + strconv.Itoa(student.ID))
	student.Age = 2 * student.ID

}

func subTaskC(student *Student) {
	// fmt.Println("subTaskC " + strconv.Itoa(student.ID))
	student.Address = "Seoul" + strconv.Itoa(student.ID)

}
