package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	fooErr := retry(foo, 4, time.Second*3)
	fmt.Printf("foo cost time: %s, err: %v\n", time.Since(start).String(), fooErr)

	start = time.Now()
	barErr := retry(bar, 4, time.Second*3)
	fmt.Printf("bar cost time: %s, err: %v\n", time.Since(start).String(), barErr)
}

// run fn in a backend goroutine
func retry(fn func(context.Context) error, count int, timeout time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	isTimeout := true
	go func() {
		for i := 0; i < count; i++ {
			if err = fn(ctx); err == nil {
				cancel()
				isTimeout = false
				return
			}
		}
	}()
	select {
	case <-ctx.Done():
		if isTimeout && err == nil {
			err = errors.New("timeout")
		}
		return
	}
}

func foo(ctx context.Context) error {
	time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond + time.Second)
	return io.EOF
}

func bar(ctx context.Context) error {
	time.Sleep(time.Minute)
	return nil
}
