package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.10.1:7777")
	if err != nil {
		fmt.Println("connect fail, err: ", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputinfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputinfo) == "Q" {
			return
		}
		_, err = conn.Write([]byte(inputinfo))
		if err != nil {
			return
		}
		buf := [1024]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err: ", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
