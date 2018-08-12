package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//
// File: main.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/08/10
//
func main() {
	// Foo_waitgroup()
	var a accumulator
	a.Foo_mutex()
	// a.Foo_atomic()

	// foo_once()
}

//===============================waitgroup===============================
func Foo_waitgroup() {
	fmt.Println("===========start Foo_waitgroup===========")
	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go foo_waitgroup(i, &wg)
	}
	wg.Wait()
	fmt.Println("===========end Foo_waitgroup===========")
}

func foo_waitgroup(a int, wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Println(a)
	wg.Done()
}

//===============================waitgroup===============================

type accumulator struct {
	mutex sync.Mutex
	wg    sync.WaitGroup
	sum   int64
}

//===============================mutex===============================

func (a *accumulator) Foo_mutex() {
	fmt.Println("===========start Foo_mutex===========")
	for i := 0; i < 16; i++ {
		a.add(i)
	}
	a.wg.Wait()
	fmt.Printf("------------sum %d--------------\n", a.sum)
	fmt.Println("===========end Foo_mutex===========")
}

func (a *accumulator) add(n int) {
	a.wg.Add(1)
	go a.foo_mutex(n)
}

func (a *accumulator) foo_mutex(n int) {
	a.mutex.Lock()
	a.sum += int64(n)
	a.mutex.Unlock()
	a.wg.Done()
}

//===============================mutex===============================

//===============================atomic===============================

func (a *accumulator) Foo_atomic() {
	a.sum = 0
	fmt.Println("===========start Foo_atomic===========")
	for i := 0; i < 16; i++ {
		a.add2(i)
	}
	a.wg.Wait()
	fmt.Printf("------------sum %d--------------\n", a.sum)
	fmt.Println("===========end Foo_atomic===========")
}

func (a *accumulator) add2(n int) {
	a.wg.Add(1)
	go a.foo_atomic(n)
}

func (a *accumulator) foo_atomic(n int) {
	atomic.AddInt64(&a.sum, int64(n))
	a.wg.Done()
}

//===============================atomic===============================

func foo_syncpool() {
	var p = sync.Pool{}
	p.New = func() interface{} {
		return []byte{}
	}
	bs := p.Get().([]byte)
	put := func(bs []byte) {
		bs = bs[:0]
		p.Put(bs)
	}
	put(bs)
}

func foo_stdmap(m map[int]struct{}) {
	m[rand.Int()%32] = struct{}{}
}

func foo_syncstdmap(m map[int]struct{}, mutex *sync.Mutex) {
	mutex.Lock()
	m[rand.Int()%32] = struct{}{}
	mutex.Unlock()
}

func foo_syncmap(m *sync.Map) {
	m.Store(rand.Int()%32, struct{}{})
}

type noCopy struct{}

func (*noCopy) Lock() {}

func foo_nocopy() {
	var c noCopy
	a := c
	_ = a
}

//================================================

func foo_once() {
	var once sync.Once
	for {
		once.Do(func() {
			fmt.Println("hello world")
		})
		fmt.Println("----------")
		time.Sleep(time.Second)
	}
}
