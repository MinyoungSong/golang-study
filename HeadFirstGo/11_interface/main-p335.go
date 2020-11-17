package main

import (
	"fmt"
)

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Beep")
}

func (r Robot) Walk() {
	fmt.Println("Powering leg")
}

type NoiseMaker interface {
	MakeSound()
}

func main() {

	var noiseMaker NoiseMaker = Robot("Robot Alex")
	noiseMaker.MakeSound()
	var robot Robot = noiseMaker.(Robot)
	robot.Walk()

	// chaining 효과로도 가능
	noiseMaker.(Robot).Walk()

}
