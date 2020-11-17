package main

import (
	"fmt"
)

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding Up")
}

func (c Car) Brake() {
	fmt.Println("Stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Car Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding Up")
}

func (t Truck) Brake() {
	fmt.Println("Stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Truck Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func main() {
	var vehicle Vehicle = Car("Toyota Yarvic")
	vehicle.Accelerate()
	vehicle.Steer("left")

	vehicle = Truck("Ford F180")
	vehicle.Brake()
	vehicle.Steer("right")

	//interface를 parameter로 받는 형태
	// doingVehicle(Car("Toyota Yarvic"), "right")
	// doingVehicle(Truck("Ford F180"), "right")
}

//interface를 parameter로 받는 형태
// func doingVehicle(v Vehicle, direction string) {
// 	v.Accelerate()
// 	v.Steer(direction)
// }
