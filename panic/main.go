//
// File: main.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/08/10
//
package main

import (
	"fmt"
	"time"
)

func main() {
	//------
	// multi_panic()
	// //-------
	// a()
	// b()
	// c()
	fmt.Println(d()) // 猜一猜输出什么
	// //------

	// multi_goroutine()
}

func multi_panic() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("e: ", e)
		}
	}()
	defer panic("panic")
	defer fmt.Println("exit")
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("e1: ", e)
		}
	}()
	panic("1")
	panic("2")
	panic("3")
	panic("4")

	fmt.Println("hello world")
}

// defer 调用时获取实参的值

func a() {
	fmt.Println("func a()")
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func b() {
	fmt.Println("func b()")
	i := 0
	defer func(i int) {
		fmt.Println("sub:", &i)
		fmt.Println(i)
	}(i)
	fmt.Println(&i)
	i++
	return
}

func c() {
	fmt.Println("func c()")
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("")
}

func d() (i int) {
	fmt.Println("func d()")
	defer func() { i++ }()
	return 3
}

func multi_goroutine() {
	go panic_goroutine()
	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Second / 2)
			fmt.Println(i)
		}
	}()
	time.Sleep(time.Minute)
}
func panic_goroutine() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover: ", e)
		}
	}()
	fmt.Println("======panic goroutine====\bsleep 1s")
	time.Sleep(time.Second)
	panic("panic sub goroutine")
}
