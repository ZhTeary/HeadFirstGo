package array

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadText(filename string) ([]float64, error) {
	numbers := make([]float64, 0) //保存命令行传进来的参数
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, err
}
