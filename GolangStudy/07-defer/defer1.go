package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

// 依次出栈
/*func main() {
	defer func1()
	defer func2()
	defer func3()
}*/
