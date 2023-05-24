package array

import "fmt"

func MyArray() {
	var A [4]string
	A[0] = "one"
	A[1] = "two"
	A[2] = "three"
	A[3] = "four"
	fmt.Println(A)
	b := [3]int{1, 2, 3}
	fmt.Println(b)
	var c [3]int = [3]int{4, 5, 6}
	fmt.Println(c)
}
