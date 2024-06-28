package main

import "fmt"

func mainMap() {

	var v int = 10
	if v == 10 {
		fmt.Println("v is 10")
	} else {
		println("v is not 10")
	}

	if a := 10; a > 4 {
		fmt.Println(a)
		// return
	}
	// for {
	// 	fmt.Println("for")

	// }
	mmap := map[int]string{
		22: "www",
		33: "sss",
		44: "ddd",
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}

	chn1 := make(chan int)
	go func() {
		chn1 <- 100
		chn1 <- 200
		chn1 <- 300
		close(chn1)
	}()
	for i := range chn1 {
		fmt.Println(i)
		// panic("panic")
	}

	m := make(map[int]int)
	m['1'] = 100
	m['2'] = 200
	v = m['1']
	fmt.Println(v)

}
