package main

import (
	"fmt"
	"greeting"
	"keyboard"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	fmt.Print("Enter your grade : ")
	garde := keyboard.GetFloat()
	fmt.Println(garde)

}
