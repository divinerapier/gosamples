//
// File: main.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/08/10
//
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// foo0()
	// bar0()
	// foo1()
	// bar1()

	//------------------------------

	// c := MakeChannel()
	// go func() {
	// 	c.send(10)
	// }()
	// fmt.Println(c.receive())

	//------------------------------
	// channel_broadcast()

	c := make(chan int)

	// consumer0
	go func() {
		select {
		case <-c:
			println("consumer0")
			return
		}
	}()
	go func() {
		select {
		case v, ok := <-c:
			fmt.Println("consumer1 ", v, ok)
			return
		}
	}()
	close(c)
	time.Sleep(time.Second)
}

// func foo0() {
// 	fmt.Println("before receiving")
// 	var c chan int
// 	<-c
// 	fmt.Println("after receiving")
// }

// func bar0() {
// 	fmt.Println("before sending")
// 	var c chan int
// 	c <- 1
// 	fmt.Println("after sending")
// }

// func foo1() {
// 	fmt.Println("before reading")
// 	c := make(chan int)
// 	<-c
// 	fmt.Println("after reading")
// }

// func bar1() {
// 	fmt.Println("before reading")
// 	c := make(chan int)
// 	c <- 1
// 	fmt.Println("after reading")
// }

// //------------------------------------------------------------------

// type Channel struct {
// 	c chan int
// }

// func MakeChannel() Channel {
// 	return Channel{
// 		c: make(chan int),
// 	}
// }

// func (c *Channel) send(n int) {
// 	c.c <- n
// }

// func (c *Channel) receive() int {
// 	return <-c.c
// }

// //---------------------------------------------------------------

// func channel_broadcast() {
// 	runtime.GOMAXPROCS(4)
// 	c := make(chan int, 100)
// 	c0 := (chan<- int)(c)
// 	c1 := (<-chan int)(c)
// 	close(c)
// 	close(c0)
// 	close(c1)
// 	go bar_consumer0(c)
// 	go bar_consumer1(c)
// 	go bar_consumer2(c)
// 	go bar_consumer3(c)
// 	foo_producer(c)
// 	fmt.Println("time.Sleep(time.Houe)")
// 	time.Sleep(time.Second * 4)
// }

// func foo_producer(c chan<- int) {
// 	var i int
// 	for {
// 		select {
// 		case c <- rand.Int() % 100:
// 			time.Sleep(time.Millisecond)
// 		}
// 		i++
// 		if i > 2000 {
// 			close(c)
// 			return
// 		}
// 	}
// }

// func bar_consumer0(c <-chan int) {
// 	caller := get_caller()
// 	for {
// 		select {
// 		case n := <-c:
// 			fmt.Printf("%s: %d\n", caller, n)
// 		}
// 	}
// 	fmt.Printf("exit %s\n", caller)
// }

// func bar_consumer1(c <-chan int) {
// 	caller := get_caller()

// 	for {
// 		select {
// 		case n, ok := <-c:
// 			if !ok {
// 				fmt.Printf("%s channel has been closed\n", caller)
// 				break
// 			}
// 			fmt.Printf("%s: %d\n", caller, n)
// 		}
// 	}
// 	fmt.Printf("exit %s\n", caller)
// }

// func bar_consumer2(c <-chan int) {
// 	caller := get_caller()
// out:
// 	for {
// 		select {
// 		case n, ok := <-c:
// 			if !ok {
// 				fmt.Println("bar_consumer2 channel has been closed")
// 				break out
// 			}
// 			fmt.Printf("%s: %d\n", caller, n)
// 		}
// 	}
// 	fmt.Printf("exit %s\n", caller)
// }

// func bar_consumer3(c <-chan int) {
// 	caller := get_caller()
// 	for n := range c {
// 		fmt.Printf("%s: %d\n", caller, n)
// 	}
// 	fmt.Printf("exit %s\n", caller)
// }

// func get_caller() string {
// 	pc, _, _, _ := runtime.Caller(1)
// 	return runtime.FuncForPC(pc).Name()
// }
