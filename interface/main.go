//
// File: main.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/08/10
//
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	// test_bytereader
	// test_readwriter()
	// test_sort()
	// test_bytereader()
	var writer io.Writer = &bytes.Buffer{}
	switch writer.(type) {
	case *os.File:
	case *bytes.Buffer:
	}
	w := writer.(*os.File)
	_ = w
}

func test_bytereader() {
	var buf bytes.Buffer
	// 为什么传递 &buf 而不是 buf
	fmt.Fprintf(&buf, "hello world: %s", "bob")
	fmt.Println(buf.String())
	fmt.Fprintf(&buf, "hello world: %s", "xxx")
	fmt.Println(buf.String())
}

func test_readwriter() {
	var writer readWriter
	fmt.Fprintf(&writer, "xxxxx")
}

func test_sort() {
	studentq := student_queue{
		student{
			ID:      0,
			Name:    "Sven",
			Stature: 175.2,
		},
		student{
			ID:      1,
			Name:    "Earth Shaker",
			Stature: 150,
		},
		student{
			ID:      2,
			Name:    "Kunkka",
			Stature: 170,
		},
		student{
			ID:      3,
			Name:    "Meppo",
			Stature: 84.9,
		},
	}
	sort.Sort(studentq)
	fmt.Println(studentq)
}
