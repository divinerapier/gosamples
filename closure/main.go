package main

import "fmt"

func main() {
	counter := getCounter()
	for i := 0; i < 10; i++ {
		fmt.Println(counter())
	}
}

func getCounter() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}
