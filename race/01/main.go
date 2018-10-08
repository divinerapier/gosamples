//
// File: main.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/10/08
//
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Printf("after %f\n", time.Since(start).Seconds())
		t.Reset(randomDuration())
	})
	time.Sleep(time.Second * 5)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
