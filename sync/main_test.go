package main

import (
	"sync"
	"testing"
)

func Benchmark_foo_stdmap(b *testing.B) {
	m := make(map[int]struct{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		foo_stdmap(m)
	}
}

func Benchmark_foo_syncmap(b *testing.B) {
	var m sync.Map
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		foo_syncmap(&m)
	}
}

func Benchmark_foo_stdmap_p(b *testing.B) {
	m := make(map[int]struct{})
	var mutex sync.Mutex
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// foo_stdmap(m)
			foo_syncstdmap(m, &mutex)
		}
	})
}

func Benchmark_foo_syncmap_p(b *testing.B) {
	var m sync.Map
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			foo_syncmap(&m)
		}
	})
}
