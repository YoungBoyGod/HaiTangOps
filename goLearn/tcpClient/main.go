package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8099")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		ReadString, err := input.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		inputInfo := strings.Trim(ReadString, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			break
		}
		_, err = conn.Write([]byte(ReadString))
		if err != nil {
			log.Fatalln(err)
		}

		var buf [512]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("recv:", string(buf[:n]))
	}
}
