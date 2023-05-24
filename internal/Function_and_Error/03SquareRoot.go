package Function_and_Error

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func SquareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, errors.New("can`t get square root of negative number")
	}
	return math.Sqrt(number), nil
}
func Square() {
	payload, err := SquareRoot(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(payload)
	}
	payload, err = SquareRoot(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(payload)
}
