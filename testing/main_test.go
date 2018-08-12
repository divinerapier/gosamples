//
// File: main_test.go
// Author: fangsihao (sihao@momenta.works)
// Created: 2018/08/12
//
package main

import "testing"

func Test_reflect_bytes_to_string(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test_reflect_bytes_to_string-00",
			args: args{
				data: []byte("hello world"),
			},
			want: "hello world",
		},
		{
			name: "Test_reflect_bytes_to_string-01",
			args: args{
				data: []byte("hello world ..."),
			},
			want: "hello world ...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reflect_bytes_to_string(tt.args.data); got != tt.want {
				t.Errorf("reflect_bytes_to_string() name = %s, got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Benchmark_normal_bytes_to_string_small(b *testing.B) {
	// do something

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = normal_bytes_to_string([]byte("hello world"))
	}
}

func Benchmark_reflect_bytes_to_string_small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reflect_bytes_to_string([]byte("hello world"))
	}
}

const largePayload = "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

func Benchmark_normal_bytes_to_string_large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = normal_bytes_to_string([]byte(largePayload))
	}
}

func Benchmark_reflect_bytes_to_string_large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reflect_bytes_to_string([]byte(largePayload))
	}
}

func Benchmark_normal_bytes_to_string_large_parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = normal_bytes_to_string([]byte(largePayload))
		}
	})
}

func Benchmark_reflect_bytes_to_string_large_parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = reflect_bytes_to_string([]byte(largePayload))
		}
	})
}
