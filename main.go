/**
*@Package:owesome_diyProtocol
*@Author: haoxiongxiao
*@Date: 2019/5/9
*@Description: create go file in owesome_diyProtocol package
 */
package main

import (
	"github.com/xhaoxiong/owesome-diyProtocol/diy"
	"net"
	"os"
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
		go diy.NewHandler(conn, 16).Do()
	}
}
