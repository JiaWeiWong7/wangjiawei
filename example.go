package main

import (
	"fmt"

	"./ftp"
)

func main() {
	Ftp, err := ftp.NewFtp("172.17.10.1:7777")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = Ftp.Login("root", "123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = Ftp.GetFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}
}
