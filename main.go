/**
*@Package:owesome_diyProtocol
*@Author: haoxiongxiao
*@Date: 2019/5/9
*@Description: create go file in owesome_diyProtocol package
 */
package main

import (
	"fmt"
	"net"
	"os"
	"owesome-diyProtocol/diy"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			os.Exit(1)
		}
		go handle2(conn)
	}
}

func handle2(conn net.Conn) {

	buffer := diy.NewBuffer(conn, 1024)

	if err := buffer.Read(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(buffer)
}