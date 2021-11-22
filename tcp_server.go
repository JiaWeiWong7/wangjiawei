package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read failed, err: ", err)
			break
		}
		recived := string(buf[:n])
		if strings.ToUpper(recived) == "Q" {
			break
		}
		fmt.Println("收到client发来的数据：", recived)
		conn.Write([]byte(recived))
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
