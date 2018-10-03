package main

import (
	"fmt"
	"sync"
)

var nc sync.WaitGroup
var count int64

func main() {

	ch := make(chan int64)

	nc.Add(2)
	go test(1, ch)
	go test(-1, ch)

	ch <- 0

	nc.Wait()
	fmt.Println(count)
}

func test(z int64, ch chan int64) {
	defer nc.Done()
	for i := 0; i < 10000; i++ {
		//count += z
		//atomic.AddInt64(&count,z)
		//runtime.Gosched()
		temp, ok := <-ch
		if !ok {
			return
		}

		temp += z

		if temp == 1 {
			close(ch)
			return
		}

		ch <- temp
	}
}

/*
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wg.Add(2)

	go player("Nadal", court)
	go player("Dj", court)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("player %s Won\n",name)
			return
		}

		n:=rand.Intn(100)
		if n%13==0{
			fmt.Printf("player %s Missed\n",name)
			close(court)
			return
		}

		fmt.Printf("Player  %s Hit %d \n",name,ball)

		ball++

		court<-ball
	}
}
*/
