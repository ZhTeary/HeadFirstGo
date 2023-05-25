package failure

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//ScanDirectory version 1: use err
//				version 2: use panic		//简化代码, 显示堆栈, 日志输出
//				version 3: use recover		//解决panic: 程序崩溃, 难看的堆栈追踪, 只想用户显示错误的信息
func ScanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			ScanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func ReportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		//恢复panic, 不是类型断言的问题仍旧panic
		panic(p)
	}
}
