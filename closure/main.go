package main

import "fmt"

func main() {
	counter0 := getCounter()
	for i := 0; i < 10; i++ {
		fmt.Println(counter0())
	}
	counter1 := getCounter()
	for i := 0; i < 10; i++ {
		fmt.Println(counter1())
	}
}

func getCounter() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}
