package main

import (
	"fmt"
	"runtime"
	"time"
	"../OtherLib"
)

const length =1024

func Test01() {
	t := time.Now()

	var arr [length * length * length]int

	for i := 0; i < length*length*length; i++ {
		arr[i] *= arr[i]
		arr[i] *= arr[i]
		arr[i] *= arr[i]
		arr[i] *= arr[i]
		arr[i] *= arr[i]
	}

	elapsed1 := time.Since(t)
	fmt.Println("A:", elapsed1)
}

func Test02() {
	t := time.Now()

	var arr [length][length][length]int

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			for k := 0; k < length; k++ {
				arr[i][j][k] *= arr[i][j][k]
				arr[i][j][k] *= arr[i][j][k]
				arr[i][j][k] *= arr[i][j][k]
				arr[i][j][k] *= arr[i][j][k]
				arr[i][j][k] *= arr[i][j][k]
			}
		}
	}

	elapsed1 := time.Since(t)
	fmt.Println("B:", elapsed1)

}

func main() {
	OtherLib.WorkPoolTest()


	return
	runtime.GOMAXPROCS(runtime.NumCPU())
	go Test01()


	time.Sleep(100*time.Second)
}
