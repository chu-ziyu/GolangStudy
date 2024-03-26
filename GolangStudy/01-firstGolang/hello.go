package main //程序的包名

import (
	"fmt"
	"time"
)

// main函数
// 语法1:花括号一定要和函数在同一行
func main() {
	//golang中的表达式，加不加分号都可以，建议不加
	fmt.Println("hello Go!")

	time.Sleep(1 * time.Second)
}
