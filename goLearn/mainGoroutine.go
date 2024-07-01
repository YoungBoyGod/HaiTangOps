package main

import (
	"fmt"
	"sync"
	"time"
)

func mai1() {
	start := time.Now().UnixMilli()
	task1()
	task2()
	end := time.Now().UnixMilli() - start
	fmt.Println("%d ", end)
}

func task1() {
	time.Sleep(time.Second * 5)
	fmt.Println("task 1")
}
func task2() {
	time.Sleep(time.Second * 3)
	fmt.Println("task 2")
}

func asyncTask1(w *sync.WaitGroup) {
	time.Sleep(time.Second * 5)
	fmt.Println("task 1")
	w.Done()
}
func asyncTask2(w *sync.WaitGroup) {
	time.Sleep(time.Second * 3)
	fmt.Println("task 2")
	w.Done()
}

func mainSync() {
	start := time.Now().UnixMilli()
	var w sync.WaitGroup
	w.Add(2)
	go asyncTask1(&w)
	go asyncTask2(&w)
	w.Wait()
	end := time.Now().UnixMilli() - start
	fmt.Println("%d ", end)
}

func hello(i int) {
	fmt.Println("hello", i)
}

func mainRoutine() {
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	fmt.Println("main go routine down")
	// time.Sleep(time.Second * 1)
}
