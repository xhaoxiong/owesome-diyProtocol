/**
*@Package:main
*@Author: haoxiongxiao
*@Date: 2019/5/9
*@Description: create go file in main package
 */
package diy

import (
	"bufio"
	"io"
)

type Handler struct {
	scanner  *bufio.Scanner
	msg      chan []byte
	DiySplit func(data []byte, atEOF bool) (advance int, token []byte, err error)
	DiyDoBuf func(buf []byte)
}

func NewHandler(conn io.Reader, len int) *Handler {
	msg := make(chan []byte, len)
	scanner := bufio.NewScanner(conn)
	return &Handler{scanner: scanner, msg: msg}
}

func (h *Handler) setDiySplit(diySplit func(data []byte, atEOF bool) (advance int, token []byte, err error)) *Handler {
	h.DiySplit = diySplit
	return h
}

func (h *Handler) setDiyDoBuf(diyDoBuf func(buf []byte)) *Handler {
	h.DiyDoBuf = diyDoBuf
	return h
}

func (h *Handler) Do() {
	h.splitBuf().doMsg()
}

func (h *Handler) splitBuf() *Handler {
	h.scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if h.DiySplit != nil {
			return h.DiySplit(data, atEOF)
		}
		return h.defaultHandleSplit(data, atEOF)
	})

	return h
}

func (h *Handler) doMsg() {
	scan := h.scanner
	for scan.Scan() {
		buf := scan.Bytes()
		if h.DiyDoBuf != nil {
			h.DiyDoBuf(buf)
		}
		h.defaultDoBuf(buf)
	}
}

func (h *Handler) defaultHandleSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	//默认的切割方式

	return
}

func (h *Handler) defaultDoBuf(buf []byte) {
	//对解析的数据的默认处理
	h.msg <- buf
}

func (h *Handler) RecvMsg() (msg []byte) {
	msg = <-h.msg
	return
}
