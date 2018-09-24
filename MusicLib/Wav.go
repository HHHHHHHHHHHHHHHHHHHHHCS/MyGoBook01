package MusicLib

import (
	"fmt"
	"time"
)

type WavPlayer struct {
	stat int
	progress int
}

func (p *WavPlayer)Play(source string){
	fmt.Println("Playing wav music",source)
	p.progress=0
	for p.progress<100{
		time.Sleep(time.Second)
		fmt.Print(".")
		p.progress+=10
	}
	fmt.Println("\nFinished playing",source)
}