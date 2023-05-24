package main

import (
	"HeadFirstGo/internal/myInterface"
)

func main() {
	var noisemakeer myInterface.NoiseMaker = myInterface.Robot("Botco Ambler")
	noisemakeer.MakeSound()
	robot := noisemakeer.(myInterface.Robot)
	robot.Walk()
}
