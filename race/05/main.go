//
// File: main.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/10/08
//
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	w.last = time.Now().UnixNano() // First conflicting access.
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(50 * time.Millisecond)
			// Second conflicting access.
			if w.last < time.Now().Add(-300*time.Millisecond).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}

func main() {
	var wd Watchdog
	wd.Start()
	for {
		r := rand.Intn(6)
		wd.KeepAlive()
		fmt.Printf("sleep %dms\n", 100*r)
		time.Sleep(time.Millisecond * 100 * time.Duration(r))
	}
}
