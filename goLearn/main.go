package main

import "fmt"

type User struct {
	Age  int
	Name string
}

func main() {
	var user User
	fmt.Printf("%v", user)
	changeUser(user)
	fmt.Printf("%v", user)
	changeUser1(&user)
	fmt.Printf("%v", user)
	var xx myInt = 100
	fmt.Println(xx.getValue())

}
func changeUser(user User) {
	user.Age = 18
}
func changeUser1(user *User) {
	user.Age = 18
}

type myInt int

func (m myInt) getValue() int {
	return int(m)
}
