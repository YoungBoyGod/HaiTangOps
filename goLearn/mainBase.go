package main

import "fmt"

// 常量
const PI float64 = 3.1415926

type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func mainBase() {
	var health int
	var level int
	fmt.Printf("health=%d, level=%d\n", health, level)
	var e = .7898789789
	fmt.Printf("e=%.5f\n", e)
	var isDead bool
	fmt.Printf("isDead=%t\n", isDead)
	var ch byte = 'A'
	var ch2 byte = 65
	fmt.Printf("ch=%c\n", ch)
	fmt.Printf("ch2=%c\n", ch2)
	str1 := "hello sdf\n"
	fmt.Println(str1)
	fmt.Printf("%d\n", len(str1))
	strByte := []byte(str1)
	fmt.Printf("%s\n", strByte)

	for i := 0; i < len(strByte); i++ {
		fmt.Printf("%d\n", strByte[i])
	}
	//在结尾追加一个Y
	strByte[len(str1)-1] = 'y'
	fmt.Printf("%s\n", strByte)

	fmt.Println(PI)
	fmt.Println(January)
	fmt.Println(February)
	var aa int = 10
	var bb int = 20
	fmt.Println(aa % bb)
	fmt.Println(aa <= bb)
	fmt.Println("++++++++++++++++")
	var a int = 10
	var p *int = &a
	// 指针指向 地址不会修改，只会在对应的地址上嵌套一层值.地址不会变。

	fmt.Println(a, p)
	fmt.Println(*p)
	*p = 20
	fmt.Println(a, p)

	fmt.Printf("打印p的地址 %v\n", p)
	fmt.Printf("打印p的值 %v\n", *p)
	fmt.Println(&p)

	fmt.Printf("=========\n")
	var cc *int
	cc = new(int)
	*cc = 100
	fmt.Println(cc)
	fmt.Println(*cc)
	fmt.Println(&cc)
}
