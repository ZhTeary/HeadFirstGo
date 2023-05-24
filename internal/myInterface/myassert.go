package myInterface

import "fmt"

func Myassert() {
	//1,2：声明接口，赋值变量
	var noiseMaker NoiseMaker = Robot("Roobot")
	//3:调用方法
	noiseMaker.MakeSound()

	//取类型断言
	robot := noiseMaker.(Robot)
	//可以调用类型(不在接口内的)函数
	robot.Walk()

	//断言失败避免异常
	isrobot, ok := noiseMaker.(Whistle)
	if ok {
		isrobot.MakeSound()
	} else {
		err := fmt.Errorf("isrobot is not Robot")
		fmt.Println(err) // == fmt.Println(err.Error())
	}
}
