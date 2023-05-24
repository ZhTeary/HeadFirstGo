package myInterface

import "fmt"

/*
	真实的error就是这样的接口
	只有一个返回string的Error()

	type error interface {
		Error() string
	}

*/

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

// OverheatError .
type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degress", o)
}

// CheckTemperature 指定函数返回原生error值
func CheckTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}
