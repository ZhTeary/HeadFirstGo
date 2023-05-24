package main

import (
	"HeadFirstGo/internal/myInterface"
	"fmt"
)

func main() {

	// error
	var err error
	err = myInterface.ComedyError("str")
	fmt.Println(err)

	err = myInterface.CheckTemperature(100, 50)
	if err != nil {
		//log.Fatal(err) // 其实是一样的 就是在Fatal里面把 自动调用了 log.Fatal(err.Error())
	}

	//stringer
	cp := myInterface.CoffeePot("LuxBrew")
	fmt.Println(cp.String())
	//很多函数会自动判断传入的参数是否满足Stringer, 如果满足就调用String方法 就和上面err和err.Error()一样
	fmt.Print(cp, "\n")
	fmt.Println(cp)
	fmt.Printf("%s\n", cp)

	fmt.Println(myInterface.Gallons(12.09248342))
	fmt.Println(myInterface.Liters(12.09248342))
	fmt.Println(myInterface.Milliliters(12.09248342))

}
