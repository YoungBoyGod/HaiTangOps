package main

import (
	"fmt"
)

func funcMain() {
	// 等同于直接赋值一个函数
	var swap func(x, y int) (int, int) //赋值之前就是nil
	swap = func(x, y int) (int, int) {
		fmt.Println(y, x)
		return y, x
	}
	fmt.Println(swap(3, 4))

	fmt.Println(swap1(3, 4))
	// 先调换 然后再相加
	r := call(swap)
	r(10)

	swap3(1, 2, 3, 4, 5, 6, 7, 8, 9, 10.10)
	var x = 1000
	var y = 2000
	changeValue(&x, &y)
	fmt.Println(x, y)
	x, y, z := yes(1, 2)
	fmt.Printf("a: %d b:%d z:%s\n", x, y, z)

	funcA := func() int {
		return 20
	}
	fmt.Println(funcA())
	func() {
		fmt.Println("匿名函数")
	}()
	generator := playerGen("testname")
	name, hp := generator()
	fmt.Println(name, hp)
}

func playerGen(name string) func() (string, int) {
	hp := 100
	return func() (string, int) {
		return name, hp
	}
}
func swap1(x, y int) (int, int) {
	return y, x
}

// 先调用输入一个函数A,然后再返回一个函数B.B再执行一个打印
func call(swap2 func(x, y int) (int, int)) func(x int) {
	y, _ := swap2(10, 20)
	return func(x int) {
		fmt.Println(y + x)
	}
}

func swap3(x int, arr ...float64) {
	fmt.Println(x)
	fmt.Println(arr)
}

func changeValue(a *int, b *int) {
	*a = 100
	*b = 200

}

func yes(x int, y int) (a int, b int, info string) {
	a = 10
	b = 20
	info = "hello"
	return
}
