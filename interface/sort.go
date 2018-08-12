//
// File: sort.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/08/12
//
package main

import (
	"encoding/json"
)

type student struct {
	ID      int64   `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Stature float64 `json:"stature,omitempty"`
}

type student_queue []student

func (s student_queue) Len() int {
	return len(s)
}

func (s student_queue) Less(i, j int) bool {
	return s[i].Stature < s[j].Stature
}

func (s student_queue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s student_queue) String() string {
	data, err := json.MarshalIndent(&s, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}
