//
// File: main.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/08/12
//
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := normal_bytes_to_string([]byte("hello world"))
	b := reflect_bytes_to_string([]byte("hello world"))
	fmt.Println(a == b)
}

func normal_bytes_to_string(data []byte) string {
	return string(data)
}

func reflect_bytes_to_string(data []byte) string {
	var stringHeader reflect.StringHeader
	sliceHeader := *(*reflect.SliceHeader)(unsafe.Pointer(&data))
	stringHeader.Data = sliceHeader.Data
	stringHeader.Len = sliceHeader.Len
	return *(*string)(unsafe.Pointer(&stringHeader))
}
