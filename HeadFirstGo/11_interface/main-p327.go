package main

import (
	"fmt"
)

type Myinterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type Mytype int

func (m Mytype) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (m Mytype) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called", f)
}

func (m Mytype) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

func (m Mytype) MethodNotInterface() {
	fmt.Println("MethodNotInterface called")
}

func main() {

	var value Myinterface
	value = Mytype(10)

	value.MethodWithoutParameters()
	value.MethodWithParameter(16)
	resultOfMethodWithReturnValue := value.MethodWithReturnValue()
	fmt.Println(resultOfMethodWithReturnValue)

	// value2는 Mytype으로 직접 지정했기 때문에 interface에 속하지 않은 메소드를 바로 사용 가능.
	// value2 := Mytype(20)
	// value2.MethodNotInterface()

}
