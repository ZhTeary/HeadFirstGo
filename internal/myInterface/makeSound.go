package myInterface

import "fmt"

type Whistle string
type Horn string
type Robot string

func (w Whistle) MakeSound() {
	fmt.Println("WWWWWWWWWWWW")
}
func (h Horn) MakeSound() {
	fmt.Println("HHHHHHHHHHHH")
}

type NoiseMaker interface {
	MakeSound()
}

func Play(n NoiseMaker) {
	n.MakeSound()
}
func (r Robot) MakeSound() {
	fmt.Println("Beep Bom")
}
func (r Robot) Walk() {
	fmt.Println("Robot is walking")
}
