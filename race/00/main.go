//
// File: main.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/10/08
//
package main

import (
	"fmt"
	"runtime"
)

func main() {
	procs := runtime.GOMAXPROCS(0)
	fmt.Println(procs)
	for i := 0; i < 1000; i++ {
		done := make(chan bool)
		m := make(map[string]string)
		m["name"] = "world"
		go func() {
			m["name"] = "data race"
			done <- true
		}()
		fmt.Println("Hello,", m["name"])
		<-done
	}
}
