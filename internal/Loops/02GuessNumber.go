package Loops

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GuessNumber() {
	//1. 获取随机数
	//2. 让用户输入
	//3. 允许猜10次 返回大了小了
	//4. 对了返回对了
	//5. 10次之后没对 就告诉答案
	//1.
	seed := time.Now().Unix()
	rand.Seed(seed)
	target := rand.Intn(100) + 1
	//2.

	//3.
	for i := 0; i < 10; i++ {
		fmt.Println("Enter your guess: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input) //回车占两个长度
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal()
		}
		if guess > target {
			fmt.Println("Bigger, you have ", 9-i, "times left")
			continue
		} else if guess < target {
			fmt.Println("Less, you have ", 9-i, "times left")
			continue
		} else {
			fmt.Println("Congratulations!!! You are Right!")
			return
		}
	}
	fmt.Println("Target is : ", target)
}
