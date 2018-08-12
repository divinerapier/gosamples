//
// File: main.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/08/13
//
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// test_array()
	test_slice()
}

func test_array() {
	fmt.Println("test_array")
	var arr0 [4]int
	fmt.Println(reflect.TypeOf(arr0))

	arr1 := [...]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(arr1))

}

func test_slice() {
	fmt.Println("test_slice")
	var arr0 [4]int
	fmt.Println(reflect.TypeOf(arr0[:]))
	arr1 := [...]int{}
	fmt.Println(reflect.TypeOf(arr1[:]))
}
