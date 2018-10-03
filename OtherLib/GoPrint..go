package OtherLib

import (
	"fmt"
	"runtime"
	"sync"
)

func GoPrint() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup

	fmt.Println("Start Goroutines")

	go func() {
		wg.Add(1)
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char <= 'z'; char ++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char <= 'Z'; char ++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting to Finish")

	wg.Wait()

	fmt.Println("End Program")
}
