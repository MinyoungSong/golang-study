package main

import (
	"fmt"
)

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

func main() {

	fmt.Println(Gallons(12.0945423523))
	fmt.Println(Liters(12.0945423523))

}
