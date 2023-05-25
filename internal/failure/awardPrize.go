package failure

import (
	"fmt"
	"math/rand"
)

/*
	实际上，调用panic并不是处理错误的理想方法
	无法访问的文件、网络故障和错误的用户输入通常应该被认为是“正常的”，
	应该通过错误值来进行适当的你的处理。
	通常，调用panic应该留给“不可能的”情况：错误表示的是程序中的错误，而不是用户方面的错误
*/
func AwardPrize() {
	doorNumber := rand.Intn(3) + 1
	switch doorNumber {
	case 1:
		fmt.Println("You win a cruise!")
	case 2:
		fmt.Println("You win a car")
	case 3:
		fmt.Println("You win a goat!")
	default:
		panic("invalid door number")
	}
}
