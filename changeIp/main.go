package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	ip, err := getIp()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ip)

}

func getIp() (string, error) {
	cmd := exec.Command("curl", "4.ipw.cn")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	output := out.String()
	return output, nil
}
