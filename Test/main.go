package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var once2 sync.Once

func p1()  {
	fmt.Println("1")
}

func p2()  {
	fmt.Println("2")
}
func main() {
	once.Do(p1)
	once.Do(p2)
	once.Do(p1)
	once.Do(p2)
	once.Do(p1)
	once.Do(p2)
	once2.Do(p2)
	once2.Do(p1)
	once2.Do(p2)
	once2.Do(p1)
	once2.Do(p2)
}
