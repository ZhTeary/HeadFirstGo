package array

func SeveralInts(numbers ...int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}
