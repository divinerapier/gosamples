package main

import (
	"fmt"
)

func main() {
	// a := map[string]struct{}{
	// 	"a": struct{}{},
	// 	"b": struct{}{},
	// 	"c": struct{}{},
	// }
	// b := map[string]struct{}{
	// 	"c": struct{}{},
	// 	"d": struct{}{},
	// 	"e": struct{}{},
	// }
	// fmt.Println(intersection(a, b))
	// fmt.Println(union(a, b))
	// fmt.Println(difference(a, b))

	// // init_map()
	// range_map()

	// m := make(map[string]int)
	// foo(m)
	// fmt.Println(m)

	foo_range()
}

func foo(m map[string]int) {
	m["a"] = 1
}

func intersection(a, b map[string]struct{}) map[string]struct{} {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	result := make(map[string]struct{})
	for k := range a {
		if _, exists := b[k]; exists {
			result[k] = struct{}{}
		}
	}
	return result
}

func union(a, b map[string]struct{}) map[string]struct{} {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	result := make(map[string]struct{}, len(a))
	for k := range a {
		result[k] = struct{}{}
	}
	for k := range b {
		result[k] = struct{}{}
	}
	return result
}

func difference(a, b map[string]struct{}) map[string]struct{} {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	result := make(map[string]struct{})
	for k := range a {
		if _, exists := b[k]; !exists {
			result[k] = struct{}{}
		}
	}
	for k := range b {
		if _, exists := a[k]; !exists {
			result[k] = struct{}{}
		}
	}
	return result
}

func init_map() {
	// 0
	var m0 = map[string]string{}
	m0["a"] = "0"
	fmt.Println(m0["a"])

	// 1
	var m1 = make(map[string]string)
	// var m1 = make(map[string]string, n)
	m1["b"] = "1"
	fmt.Println(m1["b"])

	// 2 panic
	var m2 map[string]string
	// 可以从nil map 中查找
	fmt.Println(m2["a"])
	// 不可以向nil map 赋值
	m2["c"] = "2"
	fmt.Println(m2["c"])
}

func range_map() {
	m := make(map[int]int)
	for i := 0; i < 8; i++ {
		m[i] = 1 << uint(i)
	}

	for k, v := range m {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}

	for k := range m {
		fmt.Printf("k = %v\n", k)
	}
}

type myint int

func (i myint) value_print() {
	fmt.Println(int(i))
}

func (i *myint) pointer_print() {
	fmt.Println(int(*i))
}

func foo_range() {
	m := map[int]myint{
		0: 0,
		1: 1,
		2: 2,
	}
	fmt.Println("print address")
	for k, v := range m {
		fmt.Printf("key: %p, value: %p\n", &k, &v)
	}

	fmt.Println("call value prin")
	for _, v := range m {
		v.value_print()
	}
	fmt.Println("call pointer prin")
	for _, v := range m {
		v.pointer_print()
	}
}
