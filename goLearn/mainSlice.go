package main

import (
	"fmt"
)

func mainSlice() {
	var arr [3]int

	fmt.Println(arr[0])
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	arr1 := [4]int{1, 2, 3}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	arr2 := [...]int{1, 2, 3}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	arr3 := [3]int{1: 3}
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for _, value := range arr3 {
		fmt.Println("value", value)
	}
	for i, value := range arr3 {
		fmt.Println(i, value)
	}

	var arry [4][2]int
	fmt.Println(arry)
	for k, v := range arry {
		fmt.Println(k, v)
		for i, j := range v {
			fmt.Println(i, j)
		}
	}
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2], a[2:2])
	// 修改数组需要生成新的切片
	var array [4]int
	slice := array[0:4]
	change(slice)
	fmt.Println(slice)

	var ss []int = make([]int, 2, 11)
	//长度:  容量
	fmt.Println("ss:", ss)
	fmt.Println(len(ss), cap(ss))

	vars := make([]int, 11)
	// ==vars1 := [11]int{0}
	fmt.Println(vars)
	for i := 0; i < 11; i++ {
		vars = append(vars, i)
	}
	fmt.Println(vars)

	x := 10
	for i := 0; i < x; i++ {
		qq := make([]int, i, 10)
		fmt.Printf("qq: %p\n", qq)
		fmt.Println("qq:", qq, "len:", len(qq), "cap:", cap(qq))
		qq = append(qq, 2)
		fmt.Println("qq:", qq, "len:", len(qq), "cap:", cap(qq))
		fmt.Printf("qq append %p\n", qq)

	}
	ww := make([]int, 9, 10)
	ww1 := ww[2:]
	fmt.Printf("ww v:%v,len(ww):%d,cap(ww):%d\n", ww, len(ww), cap(ww))
	fmt.Printf("ww1 v:%v,len(ww1):%d,cap(ww1):%d\n", ww1, len(ww1), cap(ww1))
	//ww v:[0 0 0 0 0 0 0 0 0],len(ww):9,cap(ww):10
	// ww1 v:[0 0 0 0 0 0 0],len(ww1):7,cap(ww1):8
	ww1[0] = 1
	// 这个修改会对wwjinx xiugai
	fmt.Printf("ww1 v:%v,len(ww1):%d,cap(ww1):%d\n", ww1, len(ww1), cap(ww1))
	ww2 := ww[3:5]
	fmt.Printf("ww2 v:%v,len(ww2):%d,cap(ww2):%d\n", ww2, len(ww2), cap(ww2))
	ww3 := ww2[1] //超过长度就会越界
	fmt.Printf("ww3 v:%v", ww3)
	ww4 := ww[0] //超过长度就会越界 ww的长度就是0
	fmt.Printf("ww4 v:%v", ww4)
	ww = append(ww, []int{100, 200, 300}...)
	fmt.Printf("ww v:%v,len(ww):%d,cap(ww):%d\n", ww, len(ww), cap(ww))

	zz := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	zz = append(zz[:2], zz[:3]...)
	fmt.Println(zz, cap(zz), len(zz))

	var valss [20]int
	for i := 0; i < 5; i++ {
		valss[i] = i * i
		fmt.Println(valss[i])
	}

	subsetLen := 5

	fmt.Println("The subset of our array has a length of:", subsetLen)
	valss[subsetLen] = 123
	subsetLen++
	fmt.Println("The subset of our array has a length of:", subsetLen)

	xx := make([]int, 5)
	fmt.Printf("%p\n", xx)
	xx = append(xx, []int{1, 2, 3, 4}...)
	fmt.Println(xx)
	fmt.Printf("%p\n", xx)
	xx = append(xx, 5)
	fmt.Println(xx)
	fmt.Printf("%p\n", xx)
	xx = append(xx, []int{1, 2, 3, 4, 5}...)
	fmt.Printf("%p\n", xx)
	xx2 := make([]int, 7)
	copy(xx2, []int{1, 2, 3, 4, 5, 6, 6, 7})
	fmt.Println(xx2)
}
func change(s []int) {
	s[0] = 1333
}
