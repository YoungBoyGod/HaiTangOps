package main

import "fmt"

type UserM struct {
	Name string
}

// 如果不使用指针，那么修改u.Name，不会修改u
func (u UserM) ModifyName(name string) {
	u.Name = name
}

func mainMethod() {
	u := UserM{}
	u.ModifyName("hello")
	fmt.Println(u)
}
