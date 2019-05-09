/**
*@Package:owesome_diyProtocol
*@Author: haoxiongxiao
*@Date: 2019/5/9
*@Description: create go file in owesome_diyProtocol package
 */
package owesome_diyProtocol

import "io"

type Buffer struct {
	reader    io.Reader
	header    []string
	buf       []byte
	bufLength int
	start     int
	end       int
}
