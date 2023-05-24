package myInterface

import "fmt"

/*
	type Anything interface{
	}
	不需要实现任何方法满足空接口,任何类型都满足它
*/

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	w, ok := thing.(Whistle)
	if ok {
		fmt.Println(w)
	}
}
