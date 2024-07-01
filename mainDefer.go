package main

import "fmt"

// func main() {
// 	x := deferTest()
// 	fmt.Println(x)

// }
// func deferTest() int {
// 	i := 1
// 	defer func() {
// 		i++
// 	}()
// 	return i
// }

func mainDefer() {
	x := deferTest()
	fmt.Println(*x)
}

// func deferTest() *int {
// 	i := 1
// 	defer func() {
// 		i++
// 	}()
// 	fmt.Println(&i, "is", i)
// 	return &i
// }

func deferTest() *int {
	i := 1
	defer func(x *int) {
		*x++
	}(&i)
	fmt.Println(&i, "is", i)
	return &i
}
