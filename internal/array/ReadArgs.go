package array

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadArgs() {
	arguments := os.Args[1:]
	var sum float64
	for _, arg := range arguments {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Println(sum)
}
