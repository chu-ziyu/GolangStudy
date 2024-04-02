package lib1

import "fmt"

// 当前lib1包提供的API
// 首字母大写表示是公共函数，否则只能在当前包内部使用
func Lib1Test() {
	fmt.Println("lib1Test()...")
}

func init() {
	fmt.Println("lib1. init() ...")
}
