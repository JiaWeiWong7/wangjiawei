package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Process(conn net.Conn) {
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputinfo := strings.Trim(input, "\r\n")
		_, err := conn.Write([]byte(inputinfo))
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
func main() {
	listen, err := net.Listen("tcp", "172.17.10.1:7777")
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			continue
		}

		go Process(conn)
	}
}
