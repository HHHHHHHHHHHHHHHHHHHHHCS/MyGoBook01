package main

import (
	"fmt"
	"sync"
	"time"
)

var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	t := time.Now()

	for i := 0; i < 1E10; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}

	elapsed1 := time.Since(t)
	t = time.Now()

	for i := 0; i < 1E10; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}

	elapsed2 := time.Since(t)

	fmt.Println("A:", elapsed1)
	fmt.Println("B:", elapsed2)
}
