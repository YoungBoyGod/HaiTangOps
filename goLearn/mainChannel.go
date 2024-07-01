package main

import "fmt"

func mainChannel() {
	ch := make(chan int)
	go recv(ch)
	ch <- 100
	fmt.Println("send ok")

	ch1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	for {
		i, ok := <-ch1
		if !ok {
			break
		}
		fmt.Println(i)
	}

}

func recv(ch chan int) {
	x := <-ch
	fmt.Println(x)
}
