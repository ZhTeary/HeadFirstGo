package myInterface

import "fmt"

type Car string
type Truck string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}
func (c Car) Brake() {
	fmt.Println("Stopping")
}
func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}
func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck) Brake() {
	fmt.Println("Stopping")
}
func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(direction string)
}
