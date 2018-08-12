//
// File: main.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/08/10
//
package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	// test_append()
	test_len()
}

func test_append() {
	var slice []int
	for i := 0; i < 16; i++ {
		slice = append(slice, i)
		fmt.Printf("value: %v, len: %d, cap: %d\n",
			slice, len(slice), cap(slice))
	}
}

func test_len() {
	var c chan int
	log.Printf("len = %d, cap = %d\n", len(c), cap(c))
	c = make(chan int)
	log.Printf("len = %d, cap = %d\n", len(c), cap(c))
	c = make(chan int, 10)
	log.Printf("len = %d, cap = %d\n", len(c), cap(c))
	c <- 1
	log.Printf("len = %d, cap = %d\n", len(c), cap(c))
	<-c
	log.Printf("len = %d, cap = %d\n", len(c), cap(c))

}
