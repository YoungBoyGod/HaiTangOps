package main

import (
	"fmt"
	"os"
	"strings"
)

func ReaderString() {
	reader := strings.NewReader("hello world xdxxx")
	p := make([]byte, 4)

	for {

		n, err := reader.Read(p)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("EOF")
				break
			}
			fmt.Println(err)

		}
		fmt.Println(n, string(p[:n]))
	}
}

func WriterString() {
	file, err := os.Create("./tst.txt")
	if err != nil {
		fmt.Println(err)
		return

	}
	defer file.Close()
	_, err = file.Write([]byte("hello world"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
func mainreader() {

}
