package ftp

import (
	"errors"
	"fmt"
	"strings"
)

func (self *Ftp) Login(user, pass string) error {
	buf := make([]byte, 1024)
	self.con.Write([]byte(fmt.Sprintf("USER is %s\r\n", user)))
	self.con.Write([]byte(fmt.Sprintf("PASS is %s\r\n", pass)))
	n, err := self.con.Read(buf)
	if err != nil {
		return err
	}
	if !strings.Contains(string(buf[:n]), "230 Logged on") {
		return errors.New(strings.TrimSpace(string(buf)))
	}
	return nil
}
