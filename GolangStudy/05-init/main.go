package main

import (
	//如果导入包但是不使用，运行时会报错，前缀加下划线可以避免
	_ "github.com/chu-ziyu/GolangStudy/05-init/lib1"
	//导入包的时候可以被包起别名
	mylib2 "github.com/chu-ziyu/GolangStudy/05-init/lib2"
)

func main() {
	//	lib1.Lib1Test()
	mylib2.Lib2Test()
}

// go mod init github.com/chu-ziyu/GolangStudy/05-init
// 使用go module模式
