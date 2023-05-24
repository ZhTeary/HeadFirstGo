package myInterface

import "fmt"

/*
	fmt包定义了fmt.Stringer 接口: 允许任何类型决定在输出时如何展示
	type Stringer interface{
		String() string
	}
	// 任何具有返回string的String方法的类型都是一个Stringer
*/

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " Coffee Pot"
}

//Demo
type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Milliliters float64

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}
