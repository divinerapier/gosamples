//
// File: main.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/10/08
//
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++
			// time.Sleep(time.Millisecond)
			fmt.Println(count)
			count--
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
}
