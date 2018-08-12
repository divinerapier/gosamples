//
// File: readwriter.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/08/10
//
package main

import (
	"fmt"
)

type readWriter struct {
	data []byte
}

// 从 interface 中读取数据到 data

func (r *readWriter) Read(p []byte) (n int, err error) {
	fmt.Println("*readWriter.Read")
	return
}

func (r *readWriter) Write(p []byte) (n int, err error) {
	fmt.Println("*readWriter.Write")
	return
}
